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
)

var (
	addr = flag.String("addr", ":1337", "TCP address to listen to")
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {
	//setup cpu core
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.WithFields(log.Fields{
		"JsonRpc": "runing",
	}).Info("runing")

	ctx, done := context.WithCancel(context.Background())
	//mq Publish
	pubTx := make(chan pubsub.Message)
	defer close(pubTx)
	go func() {
		pubsub.Publish(pubsub.Redial(ctx, *pubsub.Url), "jsonrpc.rpc", pubTx)
		done()
	}()

	//mq Subscribe
	subRx := make(chan pubsub.Message)
	defer close(subRx)
	go func() {
		pubsub.Subscribe(pubsub.Redial(ctx, *pubsub.Url), "*.rpc", subRx)
		done()
	}()

	go func() {
		data := <-subRx
		//httpserver.IdMap.Load()

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
