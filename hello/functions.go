package main

import (
    "fmt"
)

func main() {

    fmt.Println(plus(3, 4))
    fmt.Println(plusPlus(3, 4, 5))

    anInt, aString := multiReturnFunction()

    fmt.Println("anInt:", anInt, " aString:", aString)

    _, aString = multiReturnFunction() // use the blank identifier, if only a subset of the return values is going to be used

    fmt.Println("Sum of variadic function: ", sum(1, 2, 3, 4, 5))
    // or
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println("Sum of variadic function: ", sum(nums...))

    intSeq := intSeq()

    fmt.Println(intSeq())
    fmt.Println(intSeq())
    fmt.Println(intSeq())
}

func plus(a int, b int) int {
    return a + b
}

func plusPlus(a, b, c int) int { // overloading is not allowed in go
    return a + b + c
}

func multiReturnFunction() (int, string) {
    return 1, "a string"
}

func sum(nums ...int) (int) {

    sum := 0
    for _, num := range nums { // pitfall, num := range nums <- num is the index of the number within the collection
        sum += num
    }
    return sum
}

// anonymous functions as return values
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}
