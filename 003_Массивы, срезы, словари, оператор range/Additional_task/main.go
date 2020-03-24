package main

import "fmt"

func main() {
	var myArry = []int{1, 2, 3, 3, 3, 1, 1, 2, 3, 4, 4, 4, 4, 7, 7, 7, 7, 7}
	fmt.Println(myArry)
Loop:
	for i := 0; i < len(myArry); i++ {
		k := 0
		num := myArry[i]
		for j := i; j < len(myArry); j++ {
			if num == myArry[j] {
				k++
				if k > 2 {
					myArry = append(myArry[:j], myArry[j+1:]...)
					goto Loop
				}
			}
		}
	}
	fmt.Println(myArry)
}
