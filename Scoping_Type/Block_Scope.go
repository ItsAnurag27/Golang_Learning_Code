package main

import "fmt"

func main() {
    if true {
        greeting := "Hello from block scope"
        fmt.Println(greeting) // Accessible within the block
    }
    // fmt.Println(greeting) // Error: greeting is out of scope here
}