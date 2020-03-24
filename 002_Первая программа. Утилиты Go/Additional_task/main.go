package main

import (
	"Binary_Converter"
	"fmt"
)

func main() {
	var a int64 = 20
	fmt.Println(a, "=", Binary_Converter.Binary_Converter(a))
	a = 13
	fmt.Println(a, "=", Binary_Converter.Binary_Converter(a))
	a = 1
	fmt.Println(a, "=", Binary_Converter.Binary_Converter(a))
}
