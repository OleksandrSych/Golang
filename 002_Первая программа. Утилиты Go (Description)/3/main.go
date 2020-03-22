package main

import (
	"Palindrom"
	"fmt"
)

func main() {
	txt := "A palindrome is a word, number, phrase, or other sequence of characters which reads the same backward as forward, such as madam, racecar."
	fmt.Println("Original text:", txt)
	fmt.Println("Palindrom text:", Palindrom.CreatePalindrom(txt))
}
