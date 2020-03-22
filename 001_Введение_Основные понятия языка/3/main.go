package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var oper int
	var a int
	var b int
	var history string
Loop:
	fmt.Print("Выберите операцию([1] – Прибавить [2] – Отнять [3] – История [4] – Выйти): ")
	fmt.Fscan(os.Stdin, &oper)
	if oper == 3 {
		fmt.Println("\n" + history)
		fmt.Println("================================")
		goto Loop
	} else if oper == 4 {
		return
	}
	fmt.Print("\nВведите целое число а: ")
	fmt.Fscan(os.Stdin, &a)
	fmt.Print("\nВведите целое число b: ")
	fmt.Fscan(os.Stdin, &b)
	if oper == 1 {
		h := "\n" + strconv.Itoa(a) + "+" + strconv.Itoa(b) + "=" + strconv.Itoa(a+b)
		history += h
		fmt.Println(h)
	} else if oper == 2 {
		h := "\n" + strconv.Itoa(a) + "-" + strconv.Itoa(b) + "=" + strconv.Itoa(a-b)
		history += h
		fmt.Println(h)
	}
	fmt.Println("================================")
	goto Loop
}
