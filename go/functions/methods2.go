package main

import "fmt"

type Scale int

// this won't compile
/*
func (x int) squared() int {
    return x*x
}
*/

// this will work
func (x Scale) squared() Scale {
    return x*x
}

func main() {
    x := Scale(10)
    fmt.Println(x.squared())
}
