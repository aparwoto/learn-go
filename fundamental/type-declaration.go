package main

import "fmt"

func main() {
	//it is like alias

	type NOKTP string

	var KTPnumber NOKTP = "abcd123poi"

	fmt.Println(KTPnumber)

}
