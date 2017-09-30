package main

import "fmt"

func Count(ch chan int) {

	fmt.Println("Counting 1")
	ch <- 1 //发送
	fmt.Println("Counting 2")
}
func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		fmt.Println("read")
		<-ch //接收
	}
	//runtime.Gosched()
}
