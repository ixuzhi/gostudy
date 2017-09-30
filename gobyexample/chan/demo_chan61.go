package main

import "fmt"
import "time"

func send(ch chan<- int, number int) {
	fmt.Println(number)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	fmt.Println("send over")
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go send(ch1, 2)
	go send(ch2, 1)
	for {
		select {
		case v1, ok := <-ch1:
			if ok {
				fmt.Printf("%s,%d\n", "ch1", v1)
			}
		case v2, ok := <-ch2:
			if ok {
				fmt.Printf("%s,%d\n", "ch2", v2)
			}
		default:
			fmt.Println("none of the channel can process")
		}
		fmt.Println("\n")
		time.Sleep(2 * time.Second)
	}
}
