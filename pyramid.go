package main

import "fmt"

func main() {
	var rows int = 5
	for i := 1; i <= rows; i++ {
		for space := rows; space >= i; space-- {
			fmt.Print(" ")
		}
		for j := 0; j <= 2*(i-1); j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
