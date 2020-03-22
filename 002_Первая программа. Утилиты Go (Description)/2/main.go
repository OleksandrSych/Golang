package main

import (
	"NumbersToWords"
	"fmt"
)

func main() {
	num := 8
	fmt.Println(num, "  ", NTW.GetWord(num))
	num = 65
	fmt.Println(num, "  ", NTW.GetWord(num))
	num = 77
	fmt.Println(num, "  ", NTW.GetWord(num))
}
