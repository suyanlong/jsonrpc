package main

import (
	"fmt"
	"../mqserver/pubsub"
	"context"
)

func main() {
	fmt.Println("pub testing...............")
	ctx, done := context.WithCancel(context.Background())
	tx := make(chan pubsub.Message)
	defer close(tx)
	go func() {
		pubsub.Publish(pubsub.Redial(ctx, *pubsub.Url), tx)
		done()
	}()

	for i := 0; i < 100; i++ {
		str := fmt.Sprint("====", i)
		data := []byte(str)
		tx <- data
	}

	<-ctx.Done()
}
