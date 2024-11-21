package main

import "fmt"

func main() {
    sum := 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Println(sum)

    var total int
    for total = 5; total >= 0; total -= 1 {
    }
    fmt.Println(total)

    powers := 1
    for ; powers < 1000; {
        powers *= 2
    }
    fmt.Println(powers)
    
    powers2 := 1
    for powers2 < 1000 {
        powers2 *= 2
    }
    fmt.Println(powers2)

    fmt.Print("Let's G")
    for {
        fmt.Print("O")
    }
}
