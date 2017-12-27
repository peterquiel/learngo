package main


import (
	"fmt"
	"time"
)

func main () {


	i:=0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Print("\n\n")
	for j:=0; j < 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)
	}

	for true {  // endless loop
		fmt.Println("test")
		break
	}


	for { // endless loop
		fmt.Println("test")
		break
	}


	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}



	t := time.Now()
	fmt.Println(t.Date())

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}


	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}