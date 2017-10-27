package main

import (
	"encoding/json"
	"bytes"
	log "github.com/sirupsen/logrus"
	"JsonRpc/rpc"
)

func main() {
	req := bytes.NewBufferString(`{"jsonrpc": "2.0", "method": "calc_add", "params": [11, 22]}`)

	log.SetLevel(log.DebugLevel)
	var in rpc.JsonRequest
	if err := json.Unmarshal(req.Bytes(), &in); err != nil {
		log.WithFields(log.Fields{
			"req":   req.Bytes(),
			"error": err,
		}).Debugln("test error information")
	} else {
		log.WithFields(log.Fields{
			"data": in,
		}).Debugln("test information")
	}
}
