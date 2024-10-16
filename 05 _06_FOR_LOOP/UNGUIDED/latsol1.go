package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	// Inisialisasi variabel untuk menyimpan hasil penjumlahan
	sum := 0

	// Lakukan penjumlahan dari 1 hingga n
	for i := 1; i <= n; i++ {
		sum += i
	}

	// Cetak hasil penjumlahan
	fmt.Printf("Hasil penjumlahan dari 1 hingga %d adalah %d\n", n, sum)
}
