package main

import "fmt"

func main() {
	// change number below to change triangle pattern that based on number
	numb := 2
	
	stars := "*"
	for i := 1; i <= numb; i++ {
		fmt.Println(stars)
		stars += "*"
	}
}
