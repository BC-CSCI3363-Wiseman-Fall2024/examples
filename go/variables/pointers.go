package main

import "fmt"

type Pixel struct {
    x, y, z uint
}

func main() {
    value := 10
    p := &value
    var q *int = &value
    fmt.Printf("p=%v, *p=%v, &p=%v\n", p, *p, &p)
    fmt.Printf("q=%v, *q=%v, &q=%v\n", q, *q, &q)

    *q = 15
    fmt.Printf("p=%v, *p=%v, &p=%v\n", p, *p, &p)
    fmt.Printf("q=%v, *q=%v, &q=%v\n", q, *q, &q)

    pix := Pixel{1, 2, 3}
    pixp := &pix
    pixp.x = 50
    fmt.Println(pix)

}
