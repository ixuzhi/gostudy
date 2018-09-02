package main

import "fmt"
//类型 [n]T 表示拥有 n 个 T 类型的值的数组
func main(){
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0],a[1])
	fmt.Print(a)

	primes :=[6]int{2,3,5,7,11,13}
	fmt.Println(primes)
}