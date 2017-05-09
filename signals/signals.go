package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("vim-go")
	sigs := make(chan os.Signal,1)
	done := make(chan bool,1)

	signal.Notify(sigs,syscall.SIGINT,syscall.SIGTERM)

	go func() {
		sig := <- sigs
		fmt.Println("sig:",sig)

		done <- true
	}()

	fmt.Println("awaiting signal")
	<- done
	fmt.Println("exiting")
}
