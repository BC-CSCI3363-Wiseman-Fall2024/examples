package main

import "fmt"

type Player struct {
    Name string
    HP,Stamina uint
}

func main() {
    p1 := Player{"Alice", 100, 100}
    fmt.Println(p1)

    // p2.Stamina will be 0
    p2 := Player{HP: 200, Name: "Betty"}
    fmt.Println(p2)

    var p3 Player
    fmt.Println(p3)
    p3.Name = "Cindy"
    p3.HP, p3.Stamina = 50, 50
    fmt.Println(p3)
}
