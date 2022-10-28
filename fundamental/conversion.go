package main

import "fmt"

func main() {
	fmt.Println("--- int conversion ---")
	var varInt16 int16 = 128

	// it will be integer overflow because of conversion
	var varInt8 int8 = int8(varInt16)

	fmt.Println(varInt8)
	fmt.Println(varInt16)

	fmt.Println("--- string conversion ---")

	var testString = "Microservices"
	fmt.Println(testString)
	var e = testString[0]
	fmt.Println(e)
	fmt.Println(string(e))

}
