package main

import "fmt"

type Vertex struct{
	Lat,Long float64
}

var m = map[string]Vertex {
	"bell labs": Vertex{
		11,22,
	},
	"google labs": Vertex{
		33,44, 
	},
}

func main(){
	fmt.Println(m)
}