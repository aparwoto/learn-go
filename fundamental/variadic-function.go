package main

import "fmt"

func sum(number ...int) int {
	sum := 0

	for i := 0; i < len(number); i++ {
		sum += number[i]
	}

	return sum
}

func main() {
	fmt.Println(sum(-1, 2, 3, 4, 5))
}
