package main

import "fmt"

func getName() (string, string) {
	return "Alpha", "Romeo"
}

func main() {
	name1, name2 := getName()

	fmt.Println(name1, " ", name2)

	fmt.Println(name1 + " " + name2 + " " + "Juliet")

}
