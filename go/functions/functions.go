package main

import "fmt"

func bored() {
    fmt.Println("I'm sooo boooooored")
}

func callByValueTest(x int) int {
    x++
    fmt.Println("At end of function:", x)
    return x
}

func swap(x, y int) (int, int) {
    return y, x
}

func intDiv(x, y int) (quotient, remainder int) {
    quotient = x / y
    remainder = x % y
    // could also use a "naked" return:
    // return
    return quotient, remainder
}

func main() {
    bored()

    x := 1
    callByValueTest(x)
    fmt.Println("After function:", x)

    y := 3
    x, y = swap(x, y)
    fmt.Println(x, y)

    q, r := intDiv(13, 3)
    fmt.Println(q, r)
}
