package main

import "fmt"
import "time"

func send(ch chan<- int, number int) {
	fmt.Println(number)
	ch <- number
	close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go send(ch1, 2)
	go send(ch2, 1)
	for {
		select {
		case v1 := <-ch1:
			fmt.Printf("%s,%d\n", "ch1", v1)
		case v2 := <-ch2:
			fmt.Printf("%s,%d\n", "ch2", v2)
		}
		fmt.Println("end")
		time.Sleep(2 * time.Second)
	}
}
