package main

import "fmt"

func main() {
	var rows int = 5
	for i := 1; i <= rows; i++ {
		for space := 0; space <= i-1; space++ {
			fmt.Print(" ")
		}
		for j := 0; j <= 2*(rows-i); j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
