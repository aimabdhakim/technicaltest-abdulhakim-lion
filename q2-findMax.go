package main

import "fmt"

func main() {
	// change numbers below to finding max integer
	numbers := []int{1, 2, 1, 3, 1, 1}
	fmt.Println(findMax(numbers))
}

func findMax(numb []int) int {
	maxNum := 0
	for _, value := range numb {
		if value > maxNum {
			maxNum = value
		}
	}
	return maxNum
}
