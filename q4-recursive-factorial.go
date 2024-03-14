package main

import "fmt"

func main() {
	// change variable below to change factorial number
	factNumber := 7

	fmt.Println(factorial(factNumber))
}

func factorial(n int) int {
	if n == 1 {
		return n
	} else {
		return n * factorial(n-1)
	}
}
