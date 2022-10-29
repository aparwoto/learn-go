package main

import "fmt"

func getData() (name string, age, salary int) {
	name = "Anthony"
	age = 28
	salary = 3000
	return
}

func main() {
	a, b, c := getData()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
