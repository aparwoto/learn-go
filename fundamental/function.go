package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("hello world")
}

func sayAnything(s ...string) {
	fmt.Println(s)
}

func add(x int, y int) int {
	return x + y
}

func addAndMultiply(x int, y int) (resultAdd int, resultMultiply int) {
	resultAdd = x + y
	resultMultiply = x * y
	return
}

func main() {
	sayHello()
	sayHello()

	fmt.Println(add(3, 4))

	fmt.Println(addAndMultiply(3, 2))

	sayAnything("Christiano", "Ronaldo", "Manchester", "United")

}
