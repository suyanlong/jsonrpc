package main

import (
	"context"
	"fmt"

	//不能使用相对路径。
	"JsonRpc/mqserver/pubsub"
)

func main() {
	fmt.Println("sub testing...............")
	ctx, done := context.WithCancel(context.Background())
	rx := make(chan pubsub.Message)

	go func() {
		pubsub.Subscribe(pubsub.Redial(ctx, *pubsub.Url), "*.rpc", rx)
		done()
	}()

	for msg := range rx {
		fmt.Println(string(msg))
	}
	<-ctx.Done()
}
