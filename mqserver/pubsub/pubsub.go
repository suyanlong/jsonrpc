// Command pubsub is an example of a topic exchange with dynamic reliable
// membership, reading from stdin, writing to stdout.
//
// This example shows how to implement reconnect logic independent from a
// Publish/Subscribe loop with bridges to application types.

package pubsub

import (
	"context"
	"crypto/sha1"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

var Url = flag.String("Url", "amqp://guest:guest@localhost/dev", "AMQP Url for both the publisher and subscriber")

// exchange binds the publishers to the subscribers
const exchange = "CITA"

// Message is the application type for a Message.  This can contain Identity,
// or a reference to the recevier chan for further demuxing.
type Message []byte

// Session composes an amqp.Connection with an amqp.Channel
type Session struct {
	*amqp.Connection
	*amqp.Channel
}

// Close tears the connection down, taking the channel with it.
func (s Session) Close() error {
	if s.Connection == nil {
		return nil
	}
	return s.Connection.Close()
}

// Redial continually connects to the URL, exiting the program when no longer possible
func Redial(ctx context.Context, url string) chan chan Session {
	sessions := make(chan chan Session)

	go func() {
		sess := make(chan Session)
		defer close(sessions)

		for {
			select {
			case sessions <- sess:
			case <-ctx.Done():
				log.Println("shutting down Session factory")
				return
			}

			conn, err := amqp.Dial(url)
			if err != nil {
				log.Fatalf("cannot (re)dial: %v: %q", err, url)
			}

			ch, err := conn.Channel()
			if err != nil {
				log.Fatalf("cannot create channel: %v", err)
			}

			if err := ch.ExchangeDeclare(exchange, "topic", true, false, false, false, nil); err != nil {
				log.Fatalf("cannot declare topic exchange: %v", err)
			}

			select {
			case sess <- Session{conn, ch}:
			case <-ctx.Done():
				log.Println("shutting down new Session")
				return
			}
		}
	}()

	return sessions
}

// Publish publishes messages to a reconnecting Session to a topic exchange.
// It receives from the application specific source of messages.
func Publish(sessions chan chan Session, routingKey string, messages <-chan Message) {
	for session := range sessions {
		var (
			running bool
			reading = messages
			pending = make(chan Message, 1)
			confirm = make(chan amqp.Confirmation, 1)
		)

		pub := <-session

		// publisher confirms for this channel/connection
		if err := pub.Confirm(false); err != nil {
			log.Printf("publisher confirms not supported")
			close(confirm) // confirms not supported, simulate by always nacking
		} else {
			pub.NotifyPublish(confirm)
		}

		log.Printf("publishing...")

	Publish:
		for {
			var body Message
			select {
			case confirmed, ok := <-confirm:
				if !ok {
					break Publish
				}
				if !confirmed.Ack {
					log.Printf("nack Message %d, body: %q", confirmed.DeliveryTag, string(body))
				}
				reading = messages

			case body = <-pending:
				err := pub.Publish(exchange, routingKey, false, false, amqp.Publishing{
					Body: body,
				})
				// Retry failed delivery on the next Session
				if err != nil {
					pending <- body
					pub.Close()
					break Publish
				}

			case body, running = <-reading:
				// all messages consumed
				if !running {
					return
				}
				// work on pending delivery until ack'd
				pending <- body
				reading = nil
			}
		}
	}
}

// Identity returns the same host/process unique string for the lifetime of
// this process so that subscriber reconnections reuse the same queue name.
func Identity() string {
	hostname, err := os.Hostname()
	h := sha1.New()
	fmt.Fprint(h, hostname)
	fmt.Fprint(h, err)
	fmt.Fprint(h, os.Getpid())
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Subscribe consumes deliveries from an exclusive queue from a topic exchange and sends to the application specific messages chan.
func Subscribe(sessions chan chan Session, routingKey string, messages chan<- Message) {
	queue := Identity()

	for session := range sessions {
		sub := <-session

		if _, err := sub.QueueDeclare(queue, false, false, false, false, nil); err != nil {
			log.Printf("cannot consume from exclusive queue: %q, %v", queue, err)
			return
		}

		if err := sub.QueueBind(queue, routingKey, exchange, false, nil); err != nil {
			log.Printf("cannot consume without a binding to exchange: %q, %v", exchange, err)
			return
		}

		deliveries, err := sub.Consume(queue, "JsonRpc", false, false, false, false, nil)
		if err != nil {
			log.Printf("cannot consume from: %q, %v", queue, err)
			return
		}

		log.Printf("subscribed...")

		for msg := range deliveries {
			messages <- Message(msg.Body)
			sub.Ack(msg.DeliveryTag, false)
		}
	}
}
