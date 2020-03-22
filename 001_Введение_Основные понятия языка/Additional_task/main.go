package main

import "fmt"

func main() {
	fmt.Println("80F =", converter_temp(80), "C")
	fmt.Println("5F =", converter_temp(5), "C")
}

func converter_temp(tF float32) float32 {
	return (tF - 32) * 5 / 9.0
}
