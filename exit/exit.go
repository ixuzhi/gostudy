package main

import "fmt"
import "os"

func main() {
	fmt.Println("vim-go")
	defer fmt.Println("defer")
	os.Exit(3)
}
