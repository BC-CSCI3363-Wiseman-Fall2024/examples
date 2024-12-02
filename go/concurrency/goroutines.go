package main

import (
    "fmt"
    "time"
)

func task1() {
    fmt.Println("Starting task 1...")
    time.Sleep(3 * time.Second)
    fmt.Println("Finished task 1.")
}

func task2() {
    fmt.Println("Starting task 2...")
    time.Sleep(1 * time.Second)
    fmt.Println("Finished task 2.")
}

func main() {
    fmt.Println("Starting main...")
    go task1()
    go task2()
    time.Sleep(4 * time.Second)
    fmt.Println("Finished main.")
}
