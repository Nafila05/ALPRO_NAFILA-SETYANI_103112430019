package main

import "fmt"

func main() {
	var base, exponent int

	// Input bilangan dasar dan pangkat
	fmt.Print("Masukkan bilangan dasar: ")
	fmt.Scan(&base)
	fmt.Print("Masukkan bilangan pangkat: ")
	fmt.Scan(&exponent)

	// Inisialisasi hasil dengan nilai 1 (karena perkalian identitas)
	result := 1

	// Perulangan untuk menghitung pemangkatan
	for i := 0; i < exponent; i++ {
		result *= base
	}

	// Output hasil pemangkatan
	fmt.Printf("Hasil dari %d pangkat %d adalah %d\n", base, exponent, result)
}
