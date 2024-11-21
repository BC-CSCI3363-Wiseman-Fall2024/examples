package main

import "fmt"

func main() {
    // variable declarations
    var x int
    var a, b float32
    var c bool

    // initializers
    var i int = 42
    var j, k string = "hello", "world"

    // short declarations
    stuff := false
    name, valid := "alice", true

    fmt.Println(x, a, b, c)
    fmt.Println(i, j, k)
    fmt.Println(stuff, name, valid)
}

