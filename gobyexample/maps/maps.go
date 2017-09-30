package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println("map:",m)

	v1 := m["k1"]
	
	fmt.Println("map:",v1)
	fmt.Println("len:",len(m))

	delete(m,"k2")
	fmt.Println("map:",m)
	
	_,pcrs :=m["k2"]
	fmt.Println("map pcrs:",pcrs)
	n := map[string]int {"foo":1,"bar":2}
	fmt.Println("map:",n)
}
