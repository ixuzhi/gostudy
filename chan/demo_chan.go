package main

import "fmt"

var ch = make(chan int)
var content string

func set() {
	content = "It's a unbuffered channel"
	<-ch //接收
}

func main() {
	go set()
	fmt.Println("aa")
	ch <- 11 //发送后阻塞
	fmt.Println("bb")
	fmt.Println(content)
	fmt.Println("cc")
}
