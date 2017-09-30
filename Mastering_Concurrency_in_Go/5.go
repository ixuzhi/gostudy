package main

import "fmt"
import "runtime"

func showNumber(num int) {
	fmt.Println(num)
}

func listThereads() int {
	threads := runtime.GOMAXPROCS(0)
	return threads
}

func main() {
	runtime.GOMAXPROCS(3)
	fmt.Printf("%d threads available to go\n", listThereads())
	iterations := 10
	for i := 0; i < iterations; i++ {
		go showNumber(i)
	}
	//runtime.Gosched()
	fmt.Println("GoodBye!")
}
