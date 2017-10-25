package main

import "fmt"

func main() {
	count := 100
kk:
	for {
		fmt.Println(count)
		if count--; count == 0 {
			break kk;
		}
	}
}
