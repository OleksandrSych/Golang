package main

import (
	"WordFinder"
	"fmt"
)

func main() {
	var txt string = MyFinder.GetText()
	fmt.Println("'be' count -", MyFinder.Finder(txt, "be"))
	fmt.Println("'the' count -", MyFinder.Finder(txt, "the"))
}
