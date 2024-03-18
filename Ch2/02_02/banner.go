package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Hi \U0001F604"
	width := 20
	padding := (width - utf8.RuneCountInString(message)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(message)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
}
