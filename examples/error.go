package main

import "fmt"
import "strconv"

func main() {
	//error := 10;
	if vdk := 100; vdk > 10 && vdk > 99 {
		fmt.Println("ok")
	} else if vdk != 101 {
		fmt.Println("false")
	}

	fmt.Println(strconv.Itoa(2323131231231232132))
}
