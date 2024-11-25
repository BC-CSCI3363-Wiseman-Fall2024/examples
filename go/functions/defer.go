package main

import "fmt"

func f1() {
    defer fmt.Println("end of f1()")
    fmt.Println("beginning of f1()")
}

func f2() {
    defer fmt.Println("BOOM!")
    defer fmt.Println("1...")
    defer fmt.Println("2...")
    defer fmt.Println("3...")
}

func main() {
    f1()
    fmt.Println()
    f2()
}
