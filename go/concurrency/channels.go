package main

import "fmt"

func products(n1, n2 int, c chan int64) {
    var product int64 = 1
    for i := n1; i <= n2; i++ {
        product *= int64(i)
    }
    c <- product
}

// compute 20!
func main() {
    c := make(chan int64)
    go products(1, 10, c) // compute 10!
    go products(11, 20, c) // compute 20*19*..*11
    part1, part2 := <-c, <-c
    fmt.Printf("%v! = %v\n", 20, part1*part2)
}
