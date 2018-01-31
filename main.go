package main

import (
	"context"
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"jsonrpc/httpserver"
	"jsonrpc/libproto"
	"jsonrpc/mqserver/pubsub"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strconv"
)

var (
	addr     = flag.String("addr", "0.0.0.0:1337", "TCP address to listen to")
	path     = flag.String("path", "./config.json", "config information")
	capacity = flag.Int64("capacity", 50000, "channel capacity")
	///uri          = flag.String("uri", "amqp://guest:guest@localhost/dev", "AMQP URI")
	//exchange     = flag.String("exchange", "cita", "Durable AMQP exchange name")
	//exchangeType = flag.String("exchange-type", "topic", "Exchange type - direct|fanout|topic|x-custom")
	//routingKey   = flag.String("key", "*.rpc", "AMQP routing key")
)

func init() {
	flag.Parse()
	log.SetFormatter(&log.TextFormatter{})
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.ErrorLevel)
}

func main() {
	//setup cpu core

	runtime.GOMAXPROCS(runtime.NumCPU())

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	log.WithFields(log.Fields{
		"JsonRpc": "run",
	}).Info("runing")

	ctx, done := context.WithCancel(context.Background())

	//mq Publish
	//PubTx := make(chan pubsub.PubType, *capacity)
	//defer close(PubTx)
	go func() {
		pubsub.Publish(pubsub.Redial(ctx), httpserver.PubTx)
		done()
	}()

	//mq Subscribe
	var SubRx = make(chan pubsub.PubType, *capacity)
	defer close(SubRx)
	go func() {
		pubsub.Subscribe(pubsub.Redial(ctx), "*.rpc", SubRx)
		done()
	}()

	//
	go func() {
		for {
			msg := <-SubRx
			//log.Println()
			if _, typeValue, msg, err := libproto.ParseMsg(msg.Data); err == nil {
				if typeValue == libproto.MsgType_RESPONSE {
					if res, ok := msg.(*libproto.Response); ok {
						log.Info(res)
						if id, err := strconv.Atoi(string(res.RequestId)); err == nil {
							if tx, ok := httpserver.IdMap.Load(uint64(id)); ok {
								log.Info("end request", id)
								if rx, ok := tx.(chan *libproto.Response); ok {
									rx <- res
									log.Info("end request")
								}
							}
						}

					}
				}

			} else {

			}

		}
	}()

	//http server
	go func() {
		for {
			log.Info("http listen: ", *addr)
			s := fasthttp.Server{
				Handler:     httpserver.RequestHandler,
				Concurrency: 256 * 1024 * 100,
			}
			if err := s.ListenAndServe(*addr); err == nil {
				log.WithField("error", err.Error())
			}
		}
	}()

	//done()
	<-ctx.Done()
	log.Debug("end")
}
