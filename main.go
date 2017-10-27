package main

import (
	"fmt"
)

func main() {
	first := make(chan string)
	second := make(chan string)
	third := make(chan string)
	fmt.Printf("\nFirst channel:\n", first, "\n")
	fmt.Printf("\nSecond channel:\n", second, "\n")
	fmt.Printf("\nThird channel:\n", third, "\n")
}