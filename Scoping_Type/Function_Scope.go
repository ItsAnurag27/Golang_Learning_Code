package main

import "fmt"

func main() {
    message := "Hello from function scope"
    fmt.Println(message) // Accessible within main()
}

func anotherFunction() {
    // fmt.Println(message) // Error: message is not accessible here
}