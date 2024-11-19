package main

import (
	"fmt"
)

func main() {
	// Input usia
	var usia int
	fmt.Print("Masukkan usia Anda: ")
	fmt.Scan(&usia)

	// Menentukan kategori usia
	switch {
	case usia >= 0 && usia <= 12:
		fmt.Println("Kategori: Anak-anak")
	case usia >= 13 && usia <= 17:
		fmt.Println("Kategori: Remaja")
	case usia >= 18 && usia <= 64:
		fmt.Println("Kategori: Dewasa")
	case usia >= 65:
		fmt.Println("Kategori: Lansia")
	default:
		fmt.Println("Usia tidak valid")
	}
}
