package main

import "fmt"

func main() {
	var varInt int = 8
	var varString string = "This is String"
	var varBool bool = true
	var varFloat float32 = 123.123

	fmt.Println(varInt)
	fmt.Println(varString)
	fmt.Println(varBool)
	fmt.Println(varFloat)

	//declaring multiple variable

	var (
		testInt    int    = 8
		testString string = "This is another string"
		testBool   bool   = false
	)

	fmt.Println(testInt)
	fmt.Println(testString)
	fmt.Println(testBool)

}
