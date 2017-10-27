package httpserver

import (
	"github.com/valyala/fasthttp"
	log "github.com/sirupsen/logrus"
	"flag"
)

var (
	addr = flag.String("addr", ":8080", "TCP address to listen to")
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	log.Info("Request method is %q, conntinue time = %q, remoteIp = %q, ConnRequestNum = %q", ctx.Method(), ctx.ConnTime(), ctx.Host(), ctx.ConnRequestNum())
	log.Info("Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)
	log.Info(ctx.Request.Body())

	ctx.SetContentType("text/plain; charset=utf8")
	// Set arbitrary headers
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

}
