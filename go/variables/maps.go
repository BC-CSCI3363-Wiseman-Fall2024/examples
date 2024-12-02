package main

import "fmt"

func main() {
    // map literal
    ids := map[string]int{
        "Alice": 100,
        "Betty": 104,
        "Cindy": 102,
    }
    fmt.Println(ids)

    // add an element
    ids["Debra"] = 101
    fmt.Println(ids)

    // the value here will be zero
    fmt.Println(ids["Eve"])

    // check if the key exists before using the value
    id, ok := ids["Eve"]
    if !ok { 
        fmt.Println("Eve is not in the map.")
    } else {
        fmt.Println("Eve:", id)
    }
}
