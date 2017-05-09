package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
			fmt.Println("send :0")
		case ch <- 1:
			fmt.Println("send :1")
		default:
			fmt.Println("default")
		}
		i := <-ch
		fmt.Println("value received:", i)
	}
}
