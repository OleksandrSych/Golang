package main

import "fmt"

func main() {
	fmt.Println("1! =", factorial(1))
	fmt.Println("2! =", factorial(2))
	fmt.Println("3! =", factorial(3))
	fmt.Println("15! =", factorial(15))
}

func factorial(num int) (result int) {
	if num < 1 {
		result = 0
		return
	}
	result = 1
	i := 0
Loop:
	i++
	result *= i
	if i < num {
		goto Loop
	}
	return
}
