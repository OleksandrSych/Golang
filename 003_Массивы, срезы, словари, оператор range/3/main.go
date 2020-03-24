package main

import "fmt"

func main() {
	var myArray1 [15]int
	k := 1 - len(myArray1)%2
	for i := 0; i <= len(myArray1)/2; i++ {
		myArray1[len(myArray1)/2+(i-k)] = i - k
		myArray1[len(myArray1)/2-i] = i - k
	}
	fmt.Println(myArray1)
}
