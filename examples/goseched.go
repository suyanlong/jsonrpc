package main

import (
	"fmt"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func testSleep() {
	for i := 0; i < 100; i++ {
		fmt.Println(time.Now().Unix())
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println(i)
	}
}

func testGoSched() {
	for i := 0; i < 100; i++ {
		fmt.Println(time.Now().Unix())
		//time.Sleep(time.Duration(1) * time.Second)
		//runtime.Gosched()
		log.WithFields(log.Fields{
			"NumCPU":     runtime.NumCPU(),
			"Version":    runtime.Version(),
			"GoRountine": runtime.NumGoroutine(),
		}).Info("gosched")
	}
}
func main() {
	//go say("world")
	//say("hello")
	testGoSched()
	testSleep()
}
