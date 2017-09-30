package main

import "fmt"
import "time"

func main() {

	timeout := make(chan bool, 1)
	ch := make(chan int)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	select {
	case <-ch:
		fmt.Println("receive ch")
	case <-timeout:
		fmt.Println("timeout")
	}
	fmt.Println("test")
}
