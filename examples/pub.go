package main

import (
	"context"
	"fmt"

	"JsonRpc/mqserver/pubsub"
)

func main() {
	fmt.Println("pub testing...............")
	ctx, done := context.WithCancel(context.Background())
	tx := make(chan pubsub.Message)
	defer close(tx)
	go func() {
		pubsub.Publish(pubsub.Redial(ctx, *pubsub.Url), "jsonrpc.rpc", tx)
		done()
	}()

	for i := 0; i < 100; i++ {
		str := fmt.Sprint("====", i)
		data := []byte(str)
		tx <- data
	}

	//done()
	<-ctx.Done()
	fmt.Println("end")
}
