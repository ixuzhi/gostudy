package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GO runs on,%v", runtime.GOOS)

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("os:%s", os)
	}

}
