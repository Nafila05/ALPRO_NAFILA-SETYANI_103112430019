package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b) // Membaca dua input bilangan bulat dari pengguna

	result := 0

	// Menghitung hasil perkalian menggunakan penjumlahan berulang
	for i := 0; i < b; i++ {
		result += a
	}

	fmt.Println(result) // Menampilkan hasil perkalian
}
