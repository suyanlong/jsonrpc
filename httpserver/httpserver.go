package httpserver

import (
	"github.com/valyala/fasthttp"
	log "github.com/sirupsen/logrus"
	"jsonrpc/rpc"
	"fmt"
	"jsonrpc/libproto"
	"sync"
)

var IdMap sync.Map

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

	path := string(ctx.URI().Path())
	if path == "/" && ctx.IsPost() {
		//time.Sleep(time.Second * 200)
		//jsonrpc.NewServerCodec()
		if _, _, err := rpc.ParseRequest(ctx.PostBody()); err != nil {
			rx := make(chan libproto.Response, 1)
			IdMap.Store(ctx.ID(), rx)
			res := <-rx
			if res.Code == 0 {
				//

			} else {
				//TODO error
			}

			data := fmt.Sprintf("\"id\": %v,\"jsonrpc\": \"2.0\",\"result\": \"%v\"", 1, 0)
			ctx.SetBody([]byte(data))
		} else {
			ctx.SetBody([]byte("cita: this is error "))
		}

	} else {
		//http request error
		log.WithFields(log.Fields{
			"method": string(ctx.Method()),
			"path":   path,
		}).Error("http request error")

		ctx.SetBody([]byte("cita: this is error "))
	}

}
