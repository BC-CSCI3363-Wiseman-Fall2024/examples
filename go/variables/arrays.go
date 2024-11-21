package main

import "fmt"

func main() {
    var nums [5]int 
    for i := 0; i < 5; i++ {
        nums[i] = i*i
    }
    fmt.Println(nums)

    names := [5]string {
        "Alice", 
        "Betty", 
        "Cindy", 
        "Debra",
        "Eve",
    }
    fmt.Println("Full: ", names)

    slice1 := names[1:5]
    fmt.Println("slice[1:5]: ", slice1)

    // slice1[1:3] is really names[2:4]
    slice2 := slice1[1:3]
    fmt.Println("slice[2:4]: ", slice2)

    // slice2[0] is really names[2]
    slice2[0] = "Callie"
    fmt.Println("Full: ", names)
}

