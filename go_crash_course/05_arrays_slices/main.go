package main

import "fmt"

func main() {
	// Arrays
	//var fruitArr [2]string

	// assignValues
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}
	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
