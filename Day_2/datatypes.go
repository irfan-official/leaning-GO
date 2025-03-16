package main

import (
	"fmt"
)

func main() {
    var name string
    var age string
    fmt.Print("Enter your name: ")
    fmt.Scanf("%s %d", &name, &age) // Reads input until space/newline
    fmt.Println("Hello,", name)
    
}
