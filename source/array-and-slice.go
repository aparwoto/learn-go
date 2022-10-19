package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Way to initialize array
	var testArray [5]int = [5]int{1, 2, 3, 4, 5}

	// Way to initialize array
	testArray2 := [...]int{6, 7, 8, 9, 10, 11}

	// it has fixed length, so you can't append any element
	//testArray2 = append(testArray2, 12, 13, 14, 15)

	// Meanwhile, slice can be added with new element
	testSlice := []int{6, 7, 8, 9}

	testSlice = append(testSlice, 10, 11, 12, 13, 14)

	fmt.Println(reflect.ValueOf(testArray2).Kind())
	fmt.Println(reflect.ValueOf(testSlice).Kind())

	//printing elements of array
	fmt.Println(testArray)
	for i := 0; i < len(testArray); i++ {
		println(testArray[i])
	}

	fmt.Println(testArray2)
	for index, value := range testArray2 {
		fmt.Printf("index: %d , value: %d\n", index, value)
	}

	var arrayString [9]string

	var testSlice2 []string = arrayString[:]

	fmt.Println(len(arrayString))

	fmt.Println(len(testSlice2))

	//Copying array
	car := [...]string{"volvo", "bmw", "mercedez"}
	copyCar1 := car  //copy by value
	copyCar2 := &car //copy by reference

	fmt.Println(car[:])
	fmt.Println(copyCar1[:])
	fmt.Println(copyCar2[:])

	car[0] = "honda"

	fmt.Println(car[:])
	fmt.Println(copyCar1[:])
	fmt.Println(copyCar2[:])

}
