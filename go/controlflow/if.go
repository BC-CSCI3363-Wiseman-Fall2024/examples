package main

import "fmt"

func main() {
    theNumber := 16

    if theNumber > 0 {
        fmt.Println("Positive")
    } else if theNumber == 0 {
        fmt.Println("Zero")
    } else {
        fmt.Println("Negative")
    }

    if x := 10; x % 2 == 0 {
        fmt.Printf("%v is even\n", x)
    } else {
        fmt.Printf("%v is odd\n", x)
    }

    // this would be a compiler error:
    // fmt.Println(x)
}
