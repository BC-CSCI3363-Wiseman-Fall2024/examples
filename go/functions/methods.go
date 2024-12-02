package main

import "fmt"

type Rectangle struct {
    length, width float64
}

// method
func (r Rectangle) area() float64 {
    return r.length * r.width
}

// normal function equivalent
func area(r Rectangle) float64 {
    return r.length * r.width
}

func main() {
    someRectangle := Rectangle{1.5, 2.5}
    
    // use the method
    fmt.Println(someRectangle.area())

    // use the function
    fmt.Println(area(someRectangle))
}
