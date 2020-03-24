package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var myArray [50]int
	maxNum := rand.Intn(99) + 1
	for i := 0; i < len(myArray); i++ {
		myArray[i] = rand.Intn(maxNum)
	}
	if len(myArray) == 0 {
		fmt.Println("Array is empty")
	} else if len(myArray) == 1 {
		fmt.Println("My array:", myArray)
		fmt.Println("Min =", myArray[0])
	} else if len(myArray) == 2 {
		fmt.Println("My array:", myArray)
		fmt.Println("Min1 =", myArray[0])
		fmt.Println("Min2 =", myArray[1])
	} else {
		var min1 int = myArray[0]
		var min2 int = myArray[1]
		for i := 2; i < len(myArray); i++ {
			if myArray[i] < min1 || myArray[i] < min2 {
				if min1 < min2 {
					min2 = myArray[i]
				} else {
					min1 = myArray[i]
				}
			}
		}
		fmt.Println("My array:", myArray)
		fmt.Println("Min1 =", min1)
		fmt.Println("Min2 =", min2)
	}

}
