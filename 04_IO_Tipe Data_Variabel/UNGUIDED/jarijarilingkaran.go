package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	// Meminta input jari-jari dari pengguna
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&radius)

	// Menghitung luas lingkaran
	luas := math.Pi * math.Pow(radius, 2)

	// Menghitung keliling lingkaran
	keliling := 2 * math.Pi * radius

	// Mencetak hasil luas dan keliling
	fmt.Printf("Luas lingkaran: %.2f\n", luas)
	fmt.Printf("Keliling lingkaran: %.2f\n", keliling)
}
