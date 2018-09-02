package main

import "fmt"
/*
m[key] = elem
elem = m[key]
elem = m[key]

*/
func main(){
	m :=make(map[string]int)
	m["answer"] = 42
	fmt.Println("value:",m["answer"])
	m["answer"] = 43
	fmt.Println("value:",m["answer"])
	
	delete(m,"answer")
	fmt.Println("value:",m["answer"])

	v,ok := m["answer"]
	fmt.Println("value:",v,ok)

}