package main

import(
	"fmt"
	"time"
)

func main(){

	today := time.Now().Weekday()
	fmt.Println("when's %T",today,today)	
	switch today{
	case today+0:
		fmt.Println("Today")
	case today+1:
		fmt.Println("tomorrow")	
	case today+2:
		fmt.Println("in tow days")
	default:
		fmt.Println("too far away")
	}
}