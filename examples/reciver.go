package main

import "fmt"

type waiter interface {
	waiter()
}

type wait struct {
	name string
}

func (w *wait) waiter() {
	fmt.Println("", w.name);
}

type sleep struct {
	time uint
}

func (s sleep) waiter() {
	fmt.Println("", s.time)
}

func main() {
	wait := wait{name: "name"}
	wait.waiter()
	(&wait).waiter()

	sleep := sleep{time: 100}
	sleep.waiter()
	(&sleep).waiter()
}
