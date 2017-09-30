package main

import "fmt"
import "time"

func main() {
	i :=2
	switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("three")
	}

	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
			fmt.Println("It's weekday")
	default:
			fmt.Println("It's workday")
	}

	t := time.Now()
	fmt.Println("hour:",t)
	fmt.Println("year:",t.Year())
	fmt.Println("month:",t.Month())
	fmt.Println("day:",t.Day())
	fmt.Println("hour:",t.Hour())
	fmt.Println("minute:",t.Minute())
	fmt.Println("sencond:",t.Second())
	fmt.Println("nanosencond:",t.Nanosecond())
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's afternoon")
	}

	whatAmI := func(i interface{}){
		switch t := i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am integer")
		default:
			fmt.Println("Don't know type:%T",t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("string")
}
