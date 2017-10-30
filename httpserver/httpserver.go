package httpserver

import (
	"github.com/valyala/fasthttp"
	log "github.com/sirupsen/logrus"
	"jsonrpc/rpc"
	"fmt"
	"jsonrpc/libproto"
	"jsonrpc/libproto/submodules"
	"jsonrpc/libproto/topics"
	"github.com/golang/protobuf/proto"
	"sync"
	"strconv"
	"jsonrpc/mqserver/pubsub"
)

var IdMap sync.Map
var PubTx = make(chan pubsub.PubType, 10000)

func RequestHandler(ctx *fasthttp.RequestCtx) {
	log.WithFields(log.Fields{
		"ConnTime":       ctx.ConnTime(),
		"host":           string(ctx.Host()),
		"ConnRequestNum": ctx.ConnRequestNum(),
		"requestID":      ctx.ID(),
		"body":           ctx.Request.Body(),
	}).Info("request information")

	//ctx.SetContentType("text/plain; charset=utf8")
	// Set arbitrary headers
	log.Info("--------")
	path := string(ctx.URI().Path())
	if path == "/" && ctx.IsPost() {
		if values, err := rpc.ParseRequestEx(ctx.PostBody()); err == nil {
			req := libproto.Request{
				RequestId: []byte(strconv.Itoa(int(ctx.ID()))),
			}

			if values.Method == "cita_blockNumber" {
				data := libproto.Request_BlockNumber{BlockNumber: true}
				req.Req = &data

			} else if values.Method == "peerCount" {
				data := libproto.Request_Peercount{Peercount: true}
				req.Req = &data

			}

			if content, err := proto.Marshal(&req); err == nil {
				if data, err := libproto.CreateMsg(submodules.JSONRPC, topics.REQUEST, libproto.MsgType_REQUEST, content); err == nil {
					PubTx <- pubsub.PubType{
						Topic: "jsonrpc.request",
						Data:  data,
					}
				}
			}

			//TODO send mq
			rx := make(chan *libproto.Response, 1)
			log.Info("------1--")
			IdMap.Store(ctx.ID(), rx)
			res := <-rx
			log.Info("------2--")
			var result string
			if res.Code == 0 {
				//
				if res.GetBlockNumber() != 0 {
					result = fmt.Sprintf("\"id\": %s,\"jsonrpc\": \"2.0\",\"result\": \"%x\"", string(values.Id), res.GetBlockNumber())
				}

				if res.GetPeercount() != 0 {
					result = fmt.Sprintf("\"id\": %s,\"jsonrpc\": \"2.0\",\"result\": \"%x\"", string(values.Id), res.GetPeercount())
				}

			} else {
				//TODO error
				result = fmt.Sprintf("{\"id\": %s,\"jsonrpc\": \"2.0\",\"error\": \"%s\"}", string(values.Id), "error")
			}

			ctx.SetBody([]byte(result))

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
