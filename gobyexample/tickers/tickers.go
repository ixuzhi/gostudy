package main

import "fmt"
import "time"

func main() {
	fmt.Println("vim-go")
	ticker := time.NewTicker(time.Microsecond *500)
	go func(){
		for t := range ticker.C {
			fmt.Println("Tick at",t)
		}
	}()

	time.Sleep(time.Millisecond *1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
