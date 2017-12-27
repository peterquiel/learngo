package main

import (
    "fmt"
)

type displayable interface {
    display() string
}



type person struct {

    firstname string
    lastname string
}

func (p person) display() (string){ // the receiver type is person. Ponter receiver type is also possible. Just (*person) instead of
    return p.firstname + " " + p.lastname
}


func printDispalyable(d displayable) {
    fmt.Println(d.display())
}

func main() {

    fmt.Println(person{"Woody", "Allen"})

    fmt.Println(person{firstname:"Woody", lastname:"Allen"})

    fmt.Println(person{firstname:"Woody"})

    steve := person{firstname:"Steve", lastname:"Wonder"}

    fmt.Println(steve.firstname)

    stevePointer := &steve

    fmt.Println(stevePointer.lastname) // pointers to structs are automatically dereferenced

    stevePointer.lastname = "not wonder" // structs are mutable
    fmt.Println(stevePointer.lastname)


    fmt.Println(stevePointer.display())
    printDispalyable(steve)
}