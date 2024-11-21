package main

import (
    "fmt"
    "math"
)

func main() {
    someValue := 10
    // math.Sqrt() expects a float64, so have to convert someValue

    // this would be a compile time error:
    //root := math.Sqrt(someValue)

    root := math.Sqrt(float64(someValue))

    fmt.Printf("%v^0.5 = %v\n", someValue, root)
}

