package main

import (
	"fmt"
	"runtime"
	"sync"
)

var a string
var once sync.Once

func setup() {
	a = "hello world"
}

func doprint() {
	once.Do(setup)
	println(a)
}

func twoprint() {
	go doprint()
	go doprint()
}

func main() {
	fmt.Println("vim-go")
	twoprint()
	runtime.Gosched()
}
