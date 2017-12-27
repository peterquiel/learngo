package main

import (
    "errors"
    "fmt"
)

func funcOne(arg int) (int, error) {

    if arg == 42 {
        return -1, errors.New("This is not the answer")
    }
    return arg + 3, nil

}

type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob) // implementing the Error() methods gives us an own error object
}

func funcTwo(arg int) (int, error) {
    if arg == 42 {

        return -1, &argError{arg, "can't work with it"} // custom error type
    }
    return arg + 3, nil
}


func main() {
    for _, i := range []int{7, 42} {
        if r, e := funcOne(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := funcTwo(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    _, e := funcTwo(42)

    // type assertation. if e is of type *argError then ok is true and ae es of type *argError.
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
