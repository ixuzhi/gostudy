package main

import (
	"fmt"
	"time"
)

func main() {
	tick := make(chan int, 4)
	go func() {
		fmt.Println("sleep1")
		time.Sleep(2 * time.Second)
		fmt.Println("sleep2")
		count := 1
		for {
			fmt.Println("sleep3")
			time.Sleep(2 * time.Second)
			fmt.Println("sleep4")
			tick <- count
			fmt.Println("has send")
			count++
		}
	}()
	tick <- 10
	tick <- 11
	for v := range tick {
		fmt.Printf("tick: %d\n", v)
		if v == 5 {
			break
		}
	}
}
