package main

import "fmt"

/*
type Stringer interface{
	String() string
}
自描述类型
*/
type Person struct {
	Name string
	Age int
}

func (c Person) String() string{
	return fmt.Sprintf("%v,%v\n",c.Name,c.Age)
}

func main(){
	a := Person{"auther_xuxu",18}
	z := Person{"haha",21}
	fmt.Println(a,z)
}