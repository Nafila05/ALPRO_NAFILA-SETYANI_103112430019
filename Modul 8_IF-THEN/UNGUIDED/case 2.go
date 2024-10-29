package main

import (
	"fmt"
)

func main() {
	var angka int

	// Inputt
	fmt.Print("Masukkan angka untuk dicek: ")
	fmt.Scan(&angka)

	// Cek angka
	if angka%2 == 0 {
		fmt.Println("Angka adalah Genap.")
	} else {
		fmt.Println("Angka adalah Ganjil.")
	}
}
