package main

import (
	"fmt"
)

func main() {
	// Meminta input dari user
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)

	// Memastikan input adalah bilangan positif
	if n <= 0 {
		fmt.Println("Harap masukkan bilangan bulat positif.")
		return
	}

	// Mencetak segitiga bintang
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println() // Pindah baris setelah setiap baris bintang selesai
	}
}
