package main

import "fmt"

func main() {
	var n int

	// Input bilangan non-negatif
	fmt.Print("Masukkan bilangan bulat non-negatif: ")
	fmt.Scan(&n)

	// Inisialisasi hasil faktorial
	result := 1

	// Perulangan untuk menghitung faktorial
	for i := 1; i <= n; i++ {
		result *= i
	}

	// Output hasil faktorial
	fmt.Printf("Hasil faktorial dari %d adalah %d\n", n, result)
}
