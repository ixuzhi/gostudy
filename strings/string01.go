package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(strings.Contains("aabbcc", "aa"))

	fmt.Println(strings.ContainsAny("aabbcc", "dd & cc"))
}
