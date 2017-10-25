package main

import (
	"fmt"
	"../mqserver/pubsub"
	"context"
)

func main() {
	fmt.Println("sub testing...............")
	ctx, done := context.WithCancel(context.Background())
	rx := make(chan pubsub.Message)

	go func() {
		pubsub.Subscribe(pubsub.Redial(ctx, *pubsub.Url), rx)
		done()
	}()
	<-ctx.Done()
}
