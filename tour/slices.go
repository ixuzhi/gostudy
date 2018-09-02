package main
import "fmt"

// []T 切片 slice
func main(){
	primes :=[6]int{2,3,5,7,11,13}
	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Printf("%T,%T\n",primes,s)
}