package main

import "fmt"

func main() {
	person := map[string]string{
		"nama":      "arga",
		"pekerjaan": "Software Developer",
	}

	fmt.Println(person["nama"])
	fmt.Println(person["pekerjaan"])

	person["umur"] = "18"

	fmt.Println(person)

}
