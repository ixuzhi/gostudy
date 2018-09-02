package main
//Go 在不同类型的项之间赋值时需要显式转换
import(
	"fmt"
	"math"
)

func main(){
	var x,y int= 3,4
	var f float64 = math.Sqrt(float64(x*x+y*y))
	var z uint = uint(f)
	var xx uint = f
	fmt.Println(x,y,z,xx)
}