package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	c := add(2, 3)
	fmt.Printf("result is %v\n", c)
}
