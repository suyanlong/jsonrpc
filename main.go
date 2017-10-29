// This example declares a durable Exchange, and publishes a single message to
// that Exchange with a given routing key.
//
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
	addr = flag.String("addr", ":1337", "TCP address to listen to")
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

const LENGHT = 10000

func main() {
	//setup cpu core
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.WithFields(log.Fields{
		"JsonRpc": "runing",
	}).Info("runing")

	ctx, done := context.WithCancel(context.Background())
	//mq Publish
	pubTx := make(chan pubsub.Message, LENGHT)
	defer close(pubTx)
	go func() {
		pubsub.Publish(pubsub.Redial(ctx, *pubsub.Url), "jsonrpc.rpc", pubTx)
		done()
	}()

	//mq Subscribe
	subRx := make(chan pubsub.Message, LENGHT)
	defer close(subRx)
	go func() {
		pubsub.Subscribe(pubsub.Redial(ctx, *pubsub.Url), "*.rpc", subRx)
		done()
	}()

	go func() {
		data := <-subRx
		if _, typeValue, msg, err := libproto.ParseMsg(data); err == nil {
			if typeValue == libproto.MsgType_RESPONSE {
				if res, ok := msg.(*libproto.Response); ok {
					if tx, ok := httpserver.IdMap.Load(res.RequestId); ok {
						if rx, err := tx.(chan libproto.Response); err {
							rx <- *res
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
			if err := fasthttp.ListenAndServe(*addr, httpserver.RequestHandler); err == nil {
				log.WithField("error", err.Error())
			}
		}
	}()

	//done()
	<-ctx.Done()
	log.Debug("end")
}
