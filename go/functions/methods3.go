package main

import "fmt"

type Rectangle struct {
    length, width float64
}

func (r Rectangle) area() float64 {
    return r.length * r.width
}

// pointer receiver
func (r *Rectangle) transpose() {
    r.length, r.width = r.width, r.length
}

func (r Rectangle) display() {
    fmt.Printf("%v, area=%v\n", r, r.area())
}

func main() {
    r1 := Rectangle{10, 20}
    r1.display()
    r1.transpose()
    r1.display()

    rp := &Rectangle{3, 4}
    rp.display()
    rp.transpose()
    rp.display()
}
