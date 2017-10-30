package main

import (
	"context"
	"fmt"

	//不能使用相对路径。
	"jsonrpc/mqserver/pubsub"
)

func main() {
	fmt.Println("sub testing...............")
	ctx, done := context.WithCancel(context.Background())
	rx := make(chan pubsub.PubType)

	go func() {
		pubsub.Subscribe(pubsub.Redial(ctx), "*.rpc", rx)
		done()
	}()

	for msg := range rx {
		fmt.Println(msg)
	}
	<-ctx.Done()
}
