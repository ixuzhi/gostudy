package main

import "fmt"

/*
#include <stdlib.h>
#include <stdio.h>
*/
import "C"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	fmt.Println("vim-go")
	Seed(100)
	fmt.Println("Random:", Random())
}
