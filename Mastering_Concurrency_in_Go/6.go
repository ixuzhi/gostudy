package main

import (
	"fmt"
	"runtime"
	"strings"
)

var loremIpsum string
var finalIPsum string
var letterSentChan chan string

func deliverToFinal(letter string, finalIpsum *string) {
	*finalIpsum += letter
}

func capitalize(current *int, length int, letters []byte, finalIPsum *string) {
	for *current < length {
		thisLetter := strings.ToUpper(string(letters[*current]))
		deliverToFinal(thisLetter, finalIPsum)
		*current++
	}

}

func main() {
	runtime.GOMAXPROCS(2)
	index := new(int)
	*index = 0
	loremIpsum = "Awesome Go Audio and Music Authentication and OAuth Command Line Configuration Continuous Integration	CSS Preprocessors Data Structures Database 	Database Drivers Date and Time"
	letters := []byte(loremIpsum)
	length := len(letters)
	go capitalize(index, length, letters, &finalIPsum)
	/*
		go func() {
			go capitalize(index, length, letters, &finalIPsum)
		}()
	*/
	fmt.Println(length, "characters")
	fmt.Println(loremIpsum)
	fmt.Println(*index)
	fmt.Println(finalIPsum)
}
