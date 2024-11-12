package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Masukkan satu huruf: ")
	fmt.Scan(&input)

	input = strings.ToUpper(input)

	switch input {
	case "A", "I", "U", "E", "O":
		fmt.Println("Huruf Vokal")
	default:
		fmt.Println("Huruf Konsonan")
	}
}
