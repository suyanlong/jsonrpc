package httpserver

import (
	"github.com/valyala/fasthttp"
	log "github.com/sirupsen/logrus"

)

func RequestHandler(ctx *fasthttp.RequestCtx) {
	log.WithFields(log.Fields{
		"ConnTime":       ctx.ConnTime(),
		"host":           string(ctx.Host()),
		"ConnRequestNum": ctx.ConnRequestNum(),
		"requestID":      ctx.ID(),
		"body":           ctx.Request.Body(),
	}).Info("request information")

	ctx.SetContentType("text/plain; charset=utf8")
	// Set arbitrary headers
	ctx.Response.SetBody([]byte("cita: this is body "))

	path := string(ctx.URI().Path())
	if path == "/" && ctx.IsPost() {
		//time.Sleep(time.Second * 200)
		//data := string(ctx.PostBody())
		//jsonrpc.NewServerCodec()


	} else {
		//http request error
		log.WithFields(log.Fields{
			"method": string(ctx.Method()),
			"path":   path,
		}).Error("http request error")
	}

}
