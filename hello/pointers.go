package main


import (
    "fmt"
    "reflect"
)


func callByValue(anInt int) {
    anInt = 0
}

func callByReference(anIntPtr *int) { // pointer to an int
    *anIntPtr = 3 // assign a value to the dereferenced the pointer
}

func main () {
    i := 1

    fmt.Println(i)

    callByValue(i)
    fmt.Println(i)

    callByReference(&i) // &i gives us the mem address of the int ->

    fmt.Println(i)
    fmt.Println(&i) // pointer can be printed
    fmt.Println(reflect.TypeOf(&i)) // type of an pointer to int is *int
}
