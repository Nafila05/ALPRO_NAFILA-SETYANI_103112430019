package main

import (
	"fmt"
)

func main() {
	// Meminta input dari user
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	// Memastikan input adalah bilangan positif
	if n <= 0 {
		fmt.Println("Harap masukkan bilangan bulat positif.")
		return
	}

	// Menghitung jumlah deret dari 1 hingga n
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Printf("Jumlah dari 1 hingga %d adalah: %d\n", n, sum)
}
