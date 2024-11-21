package main

import "fmt"

func main() {
    orig := []string{"a", "b", "c", "d", "e", "f"}
    fmt.Printf("%v, len:%v, cap:%v\n", orig, len(orig), cap(orig))

    // empty slice
    s := orig[1:1]
    fmt.Printf("s[1:1]: %v, len:%v, cap:%v\n", s, len(s), cap(s))

    // extend the length, still within the capacity
    s = s[:4]
    fmt.Printf("s[:4]: %v, len:%v, cap:%v\n", s, len(s), cap(s))

    // loop over the slice
    for i, val := range s {
        fmt.Printf("s[%v] = %v, ", i, val)
    }
    fmt.Println()

    // move the starting point up by one element
    s = s[1:]
    fmt.Printf("s[1:]: %v, len:%v, cap:%v\n", s, len(s), cap(s))

    // loop over the indices only
    for i := range s {
        fmt.Printf("s[%v] = %v, ", i, s[i])
    }
    fmt.Println()

    // move the starting point up again
    s = s[3:]
    fmt.Printf("s[3:]: %v, len:%v, cap:%v\n", s, len(s), cap(s))

    // this will be a run-time error:
    s = s[:2]
}
