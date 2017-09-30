package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

var initialString string
var finalString string
var stringLength int

func addToFinalStack(letterChannel chan string, wg *sync.WaitGroup) {
	letter := <-letterChannel
	finalString += letter
	wg.Done()
}

func capitalize(letterChannel chan string, currentLetter string, wg *sync.WaitGroup) {
	thisLetter := strings.ToUpper(currentLetter)
	wg.Done()
	letterChannel <- thisLetter
}

func main() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup

	initialString = "Awesome Go Audio and Music Authentication and OAuth Command Line Configuration Continuous Integration	CSS Preprocessors Data Structures Database 	Database Drivers Date and Time"
	initialBytes := []byte(initialString)
	var letterChannel chan string = make(chan string)
	stringLength = len(initialBytes)
	for i := 0; i < stringLength; i++ {
		wg.Add(2)
		go capitalize(letterChannel, string(initialBytes[i]), &wg)
		go addToFinalStack(letterChannel, &wg)
		wg.Wait()
	}
	fmt.Println(initialString)
	fmt.Println(finalString)
}
