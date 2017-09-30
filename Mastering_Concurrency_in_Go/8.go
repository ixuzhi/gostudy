package main

import (
	"fmt"
	"strings"
	"sync"
)

var initialString string
var initialBytes []byte
var stringLength int
var finalString string
var letterProcessed int
var applicationStatus bool
var wg sync.WaitGroup

func getLetters(gQ chan string) {
	for i := range initialBytes {
		gQ <- string(initialBytes[i])
	}
}

func capitalizeLetters(gQ chan string, sQ chan string) {
	for {
		fmt.Println(letterProcessed)
		if letterProcessed >= stringLength {
			applicationStatus = false
			break
		}
		select {
		case letter := <-gQ:
			capitalLetter := strings.ToUpper(letter)
			finalString += capitalLetter
			letterProcessed++
		}
	}
}

func main() {
	applicationStatus = true
	getQueue := make(chan string)
	stackQueue := make(chan string)

	initialString = "Awesome Go Audio and Music Authentication and OAuth Command Line Configuration Continuous Integration	CSS Preprocessors Data Structures Database 	Database Drivers Date and Time"
	initialBytes := []byte(initialString)
	stringLength = len(initialBytes)
	letterProcessed = 0
	fmt.Println("let start capitalizing")
	go getLetters(getQueue)
	capitalizeLetters(getQueue, stackQueue)

	for {
		if applicationStatus == false {
			fmt.Println("Done")
			fmt.Println(finalString)
			break
		}
	}
	close(getQueue)
	close(stackQueue)

}
