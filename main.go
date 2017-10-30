package main

import (
	"context"
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"jsonrpc/httpserver"
	"jsonrpc/mqserver/pubsub"
	"runtime"
	"jsonrpc/libproto"
)

var (
	addr     = flag.String("addr", "0.0.0.0:1337", "TCP address to listen to")
	path     = flag.String("path", "./config.json", "config information")
	capacity = flag.Int64("capacity", 10000, "channel capacity")
	///uri          = flag.String("uri", "amqp://guest:guest@localhost/dev", "AMQP URI")
	//exchange     = flag.String("exchange", "cita", "Durable AMQP exchange name")
	//exchangeType = flag.String("exchange-type", "topic", "Exchange type - direct|fanout|topic|x-custom")
	//routingKey   = flag.String("key", "*.rpc", "AMQP routing key")
)

func init() {
	flag.Parse()
	log.SetFormatter(&log.TextFormatter{})
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {
	//setup cpu core
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.WithFields(log.Fields{
		"JsonRpc": "run",
	}).Info("runing")

	ctx, done := context.WithCancel(context.Background())
	//mq Publish
	pubTx := make(chan pubsub.PubType, *capacity)
	defer close(pubTx)
	go func() {
		pubsub.Publish(pubsub.Redial(ctx), pubTx)
		done()
	}()

	//mq Subscribe
	subRx := make(chan pubsub.PubType, *capacity)
	defer close(subRx)
	go func() {
		pubsub.Subscribe(pubsub.Redial(ctx), "*.rpc", subRx)
		done()
	}()

	//
	go func() {
		data := <-subRx
		if _, typeValue, msg, err := libproto.ParseMsg(data); err == nil {
			if typeValue == libproto.MsgType_RESPONSE {
				if res, ok := msg.(*libproto.Response); ok {
					if tx, ok := httpserver.IdMap.Load(res.RequestId); ok {
						if rx, err := tx.(chan *libproto.Response); err {
							rx <- res
						}
					}
				}
			}

		} else {

		}
	}()

	//http server
	go func() {
		for {
			log.Info("http listen: ", *addr)
			if err := fasthttp.ListenAndServe(*addr, httpserver.RequestHandler); err == nil {
				log.WithField("error", err.Error())
			}
		}
	}()

	//done()
	<-ctx.Done()
	log.Debug("end")
}
