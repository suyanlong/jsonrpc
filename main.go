// This example declares a durable Exchange, and publishes a single message to
// that Exchange with a given routing key.
//
package main

import (
	"github.com/valyala/fasthttp"
	log "github.com/sirupsen/logrus"
	"flag"
	"jsonrpc/httpserver"
)

var (
	addr = flag.String("addr", ":8080", "TCP address to listen to")
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.WithFields(log.Fields{
		"JsonRpc": "runing",
	}).Info("runing")

	//http server
	for {
		if err := fasthttp.ListenAndServe(*addr, httpserver.RequestHandler); err == nil {
			log.WithField("error", err.Error())
		}
	}

}
