package main

import (
	"math"
	"fmt"
)

const constantString = "this is const string"

func main() {

	fmt.Println(constantString)

	const n = 500000
	const e = 10e20

	fmt.Println(e / n )

	fmt.Println(int64(e/n))

	fmt.Print(math.Sqrt2)
}
