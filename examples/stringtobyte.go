package main

import "fmt"

func main() {
	//类型转换
	str := string("123你好")
	data := []byte(str)
	fmt.Println(str)
	fmt.Println(data)
}
