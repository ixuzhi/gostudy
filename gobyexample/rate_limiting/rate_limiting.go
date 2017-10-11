package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		time.Sleep(time.Second * 1)
		fmt.Println("requests", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
		time.Sleep(time.Second * 1)
	}
	close(burstyLimiter)
	//如果不close range 不会结束

	for t1 := range burstyLimiter {
		fmt.Println("time.Now()", t1)
	}
	fmt.Println("aa\n")

	burst := make(chan time.Time, 3)
	func() {
		for t := range time.Tick(time.Second * 2) {
			fmt.Println("t:", t)
			burst <- t
			fmt.Println("burstyLimiter")
		}
	}()
	close(burst)
	for t := range burst {
		fmt.Println(t)
	}
	fmt.Println("bb")
	/*
		burstyRequests := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			burstyRequests <- i
		}
		close(burstyRequests)

		for req := range burstyRequests {
			<-burstyRequests
			fmt.Println("request ", req, time.Now())
		}
	*/
}
