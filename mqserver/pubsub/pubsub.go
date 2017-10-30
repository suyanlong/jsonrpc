package pubsub

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
	"flag"
)

var (
	uri          = flag.String("uri", "amqp://guest:guest@localhost/dev", "AMQP URI for both the publisher and subscriber")
	exchange     = flag.String("exchange", "cita", "Durable AMQP exchange name")
	exchangeType = flag.String("exchange-type", "topic", "Exchange type - direct|fanout|Topic|x-custom")
	queue        = flag.String("queue", "jsonrpc", "subscriber queue")
	routingKey   = flag.String("key", "*.rpc", "subscriber routing key")
)

func init() {
	flag.Parse()
}

// Message is the application type for a Message.  This can contain Identity,
// or a reference to the recevier chan for further demuxing.
//type Message []byte

type PubType struct {
	Topic string
	Data  []byte
}

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
func Redial(ctx context.Context) chan chan Session {
	sessions := make(chan chan Session)

	go func() {
		sess := make(chan Session)
		defer close(sessions)

		for {
			//
			select {
			case sessions <- sess:
			case <-ctx.Done():
				log.Println("shutting down Session factory")
				return
			}

			//connection rabbitmq
			conn, err := amqp.Dial(*uri)
			if err != nil {
				log.Fatalf("cannot (re)dial: %v: %q", err, *uri)
			}

			ch, err := conn.Channel()
			if err != nil {
				log.Fatalf("cannot create channel: %v", err)
			}

			if err := ch.ExchangeDeclare(*exchange, *exchangeType, true, false, false, false, nil); err != nil {
				log.Fatalf("cannot declare Topic exchange: %v", err)
			}

			//
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

// Publish publishes messages to a reconnecting Session to a Topic exchange.
// It receives from the application specific source of messages.
func Publish(sessions chan chan Session, messages <-chan PubType) {
	for session := range sessions {
		var (
			running bool
			reading = messages
			pending = make(chan PubType, 1)
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
			var body PubType
			select {
			case confirmed, ok := <-confirm:
				if !ok {
					break Publish
				}
				if !confirmed.Ack {
					log.Printf("nack Message %d, body: %q", confirmed.DeliveryTag, string(body.Data))
				}
				reading = messages

			case body = <-pending:
				err := pub.Publish(*exchange, body.Topic, false, false, amqp.Publishing{
					Body: body.Data,
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

// Subscribe consumes deliveries from an exclusive queue from a Topic exchange and sends to the application specific messages chan.
func Subscribe(sessions chan chan Session, routingKey string, messages chan<- PubType) {
	//queue := Identity()

	for session := range sessions {
		sub := <-session

		if _, err := sub.QueueDeclare(*queue, true, false, false, false, nil); err != nil {
			log.Printf("cannot consume from exclusive queue: %q, %v", queue, err)
			return
		}

		if err := sub.QueueBind(*queue, routingKey, *exchange, false, nil); err != nil {
			log.Printf("cannot consume without a binding to exchange: %q, %v", *exchange, err)
			return
		}

		if deliveries, err := sub.Consume(*queue, "JsonRpc", false, false, false, false, nil); err != nil {
			log.Printf("cannot consume from: %q, %v", queue, err)
			return
		} else {

			log.Printf("subscribed...")
			for msg := range deliveries {
				messages <- PubType{
					Data:  msg.Body,
					Topic: msg.RoutingKey,
				}
				sub.Ack(msg.DeliveryTag, true)
			}
		}
	}
}
