package main

import (
	"fmt"
)

func Pipeline(from chan string, to chan string, n int) {
	i := 0
	for i < n {
		select {
		case value := <-from:
			to <- value
			i = i + 1
		}
	}
}

func main() {

	// Empty channels
	first := make(chan string)
	second := make(chan string)
	third := make(chan string)

	// Write to the first channel
	go func() {
		first <- "First value"
	}()
	go func() {
		first <- "Second value"
	}()
	go func() {
		first <- "Third value"
	}()

	// Pass data from the first channel into the second channel and
	// from the second channel into the third channel.
	go func() {
		Pipeline(first, second, 3)
	}()
	go func() {
		Pipeline(second, third, 3)
	}()

	// Render data from the third channel
	i := 0
	for i < 3 {
		select {
		case value := <-third:
			fmt.Printf(value + "\n")
			i = i + 1
		}
	}
}
