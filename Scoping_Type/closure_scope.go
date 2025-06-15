package main

import "fmt"

func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := createCounter()
	fmt.Println(counter()) // Prints 1
	fmt.Println(counter()) // Prints 2
}
