package main

import "fmt"

func main() {
	var N int = 10
	var a, b int = 0, 1
	for i := 0; i < N; i++ {
		fmt.Println(a)
		Next := a + b
		a = b
		b = Next
	}
}
