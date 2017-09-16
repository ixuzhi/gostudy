package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.1, 30.2,
	},
	"Google": Vertex{
		40.2, 30.3,
	},
}

func main() {
	fmt.Println(m)
}
