package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var myArray [15]int
	for i := 0; i < len(myArray); i++ {
		myArray[i] = rand.Intn(100)
	}
	var invertArray [15]int
	for i := 0; i < len(invertArray); i++ {
		invertArray[i] = myArray[len(invertArray)-1-i]
	}
	fmt.Println("My array    :", myArray)
	fmt.Println("Invert array:", invertArray)
}
