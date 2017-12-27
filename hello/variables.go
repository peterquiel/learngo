package main

import "fmt"

func main() {

	var a = "initialized" // type derived from value

	fmt.Println(a)

	var b,c =  2 , 3 // multi variable initialization

	println(b + c )

	fmt.Println(b + c ) // cast to string

	var e int // ok, implicit 0 init, like java does

	fmt.Println(e)

	var f bool // implicit false init

	fmt.Println(f)

	g := "short for var g string ='xyz' "

	fmt.Println(g)

}
