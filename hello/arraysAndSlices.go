package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a [5] int // arrays are initialized with zeros
	fmt.Println(a)

	a[4] = 100
	fmt.Println("complete array:", a)
	fmt.Println("get:", a[4])
	fmt.Println("array length", len(a))

	b := [5]int{1, 2, 3, 4, 5} // inline init
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// slices
	s := make([]string, 3)
	fmt.Println("emp:", s, " type:", " type:", reflect.TypeOf(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:len(s)]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t, " type:", reflect.TypeOf(t)) // array vs. slice type is [3]int vs []int
}
