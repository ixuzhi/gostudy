package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.1, 30.2},
	"Google":    {30.4, 50.3},
}

func main() {
	fmt.Println(m)
}
