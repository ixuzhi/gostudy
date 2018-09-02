package main

import "fmt"

type Vertex struct{
	Lat,Long float64
}

var m = map[string]Vertex{
	"bell labs":{1,2},
	"google labs":{3,4},
}

func main(){
	fmt.Println(m)
}