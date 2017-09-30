package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job, goGroup *sync.WaitGroup) {
	filename := j.text + ".txt"
	filecontent := ""
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		filecontent += j.text
		fmt.Println(j.text)
		j.i++
	}
	err := ioutil.WriteFile(filename, []byte(filecontent), 0644)
	if err != nil {
		panic("something went away")
	}
}

func main() {
	hello := new(Job)
	world := new(Job)
	goGroup := new(sync.WaitGroup)
	hello.text = "hello"
	hello.i = 0
	hello.max = 3

	world.text = "world"
	world.i = 0
	world.max = 5

	go outputText(hello, goGroup)
	go outputText(world, goGroup)
	//runtime.Gosched()
}
