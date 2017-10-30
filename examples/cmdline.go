package main

import (
	"flag"
	"fmt"
)

var (
	capacity = flag.Int64("capacity", 10000, "channel capacity")
	path     = flag.String("path", "./config.json", "config information")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Println(flag.Args())
	fmt.Println(*path)
}
