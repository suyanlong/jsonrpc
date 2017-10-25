package main

import "fmt"

func main() {
	//error := 10;
	if vdk := 100; vdk > 10 && vdk > 99 {
		fmt.Println("ok")
	} else if vdk != 101 {
		fmt.Println("false")
	}
}
