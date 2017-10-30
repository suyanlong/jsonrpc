package main

import (
	"context"
	"fmt"
	"jsonrpc/mqserver/pubsub"
)

func main() {
	fmt.Println("pub testing...............")
	ctx, done := context.WithCancel(context.Background())
	tx := make(chan pubsub.PubType)
	defer close(tx)
	go func() {
		pubsub.Publish(pubsub.Redial(ctx), tx)
		done()
	}()

	for i := 0; i < 100; i++ {
		str := fmt.Sprint("====", i)
		data := []byte(str)
		tx <- pubsub.PubType{
			Topic: "jsonrpc.rpc",
			Data:  data,
		}
	}

	//done()
	<-ctx.Done()
	fmt.Println("end")
}
