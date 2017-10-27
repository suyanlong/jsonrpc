// This example declares a durable Exchange, and publishes a single message to
// that Exchange with a given routing key.
//
package main

import (
	"github.com/valyala/fasthttp"
	log "github.com/sirupsen/logrus"
	"flag"
	"net/rpc/jsonrpc"
)

var (
	addr = flag.String("addr", ":8080", "TCP address to listen to")
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	log.Info(log.WithField("time", 12))
	log.Info("Request method is %s, conntinue time = %v, remoteIp = %s, ConnRequestNum = %d ", string(ctx.Method()), ctx.ConnTime(), string(ctx.Host()), ctx.ConnRequestNum())
	log.Info("Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)
	log.Info(ctx.Request.Body())

	ctx.SetContentType("text/plain; charset=utf8")
	// Set arbitrary headers
	ctx.Response.SetBody([]byte("123"))

}

func init() {
	format := log.TextFormatter{}
	log.SetFormatter(&format)
	//log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	log.WithFields(log.Fields{
		"JsonRpc": "runing",
	}).Info("runing")

	fasthttp.ListenAndServe(*addr, requestHandler)
	fasthttp.AcquireResponse()

}
