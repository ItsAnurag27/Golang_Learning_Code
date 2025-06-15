package main

import (
	"errors"
	"fmt"
)

func process() error {
	fmt.Println("starting pracess")
	err := errors.New("an error occured")
	if err != nil {
		return fmt.Errorf("error: %w", err)

	}
	fmt.Println("process completed success")
	return nil
}
func main() {
	err := process()
	if err != nil {
		fmt.Println("process failed:", err)
	}
}
