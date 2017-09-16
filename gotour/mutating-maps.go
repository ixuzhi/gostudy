package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The answer:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("the answer:", m["Answer"])

	delete(m, "Answer")

	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("the value :", v, "present", ok)

}
