package main

import "fmt"

func main() {
	a := 2
	b := 3
	display_variables(a, b)
	fmt.Println("Change")
	a, b = b, a
	display_variables(a, b)
}

func display_variables(a int, b int) {
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
