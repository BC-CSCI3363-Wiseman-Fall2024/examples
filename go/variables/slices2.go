package main

import "fmt"

func main() {
    s := make([]int, 3)
    fmt.Printf("s:%v, len:%v, cap:%v\n", s, len(s), cap(s))
    for i := range s {
        s[i] = i + 1
    }
    fmt.Printf("s:%v, len:%v, cap:%v\n", s, len(s), cap(s))

    s = append(s, 4)
    s = append(s, 5, 6)
    fmt.Printf("s:%v, len:%v, cap:%v\n", s, len(s), cap(s))

    t := make([]int, 2, 4)
    t[0] = 11
    t[1] = 12
    fmt.Printf("t:%v, len:%v, cap:%v\n", t, len(t), cap(t))

    t = append(t, 0)
    fmt.Printf("t:%v, len:%v, cap:%v\n", t, len(t), cap(t))
    
    t = t[:cap(t)]
    fmt.Printf("t:%v, len:%v, cap:%v\n", t, len(t), cap(t))

    u := t[1:3]
    fmt.Printf("u:%v, len:%v, cap:%v\n", u, len(u), cap(u))

    u = append(u, 42)
    fmt.Printf("u:%v, len:%v, cap:%v\n", u, len(u), cap(u))
    fmt.Printf("t:%v, len:%v, cap:%v\n", t, len(t), cap(t))
}
