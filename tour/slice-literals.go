package main

import "fmt"
// [x] []数组、切片文法
func main(){
	q := []int{2,3,5,7,11,13}
	fmt.Println(q)
	r := []bool{true,false,true,false,true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2,true},
		{3,false},
		{5,true},
		{7,false},
		{11,false},
	}
	fmt.Println(s)
}