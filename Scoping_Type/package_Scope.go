package mypkg

import "fmt"

// Package-scoped variables
var (
    Greeting = "Hello from package scope" // Exported (accessible outside the package)
    secret   = "This is private"          // Not exported
)

// PrintGreeting prints the greeting message.
func PrintGreeting() {
    fmt.Println(Greeting) // Accessible within the package
}