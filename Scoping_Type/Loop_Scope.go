package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i) // i is accessible within the loop
    }
    // fmt.Println(i) // Error: i is out of scope here
}