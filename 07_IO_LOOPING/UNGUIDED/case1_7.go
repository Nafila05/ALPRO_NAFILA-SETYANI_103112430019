package main

import (
	"fmt"
)

func main() {
	// Meminta input dari pengguna
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	// Loop untuk mencetak hasil kuadrat dari 1 hingga N
	fmt.Println("Hasil kuadrat dari 1 hingga", N, "adalah:")
	for i := 1; i <= N; i++ {
		fmt.Println(i, "kuadrat =", i*i)
	}
}
