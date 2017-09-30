package main

import (
	"fmt"
	"time"
)

func main() {
	tick := make(chan int, 1)
	go func() {
		time.Sleep(2 * time.Second)
		count := 1
		for {
			time.Sleep(1 * time.Second)
			tick <- count
			if count == 5 {
				close(tick)
			}
			count++
			tick <- count
		}
	}()

	for v := range tick {
		fmt.Printf("%d\n", v)
	}
}
