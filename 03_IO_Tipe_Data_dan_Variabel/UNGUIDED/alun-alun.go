package main

import "fmt"

func main() {
	// Panjang sisi alun-alun
	var sisi float64 = 27

	// Menghitung keliling (perimeter)
	keliling := 4 * sisi

	// Menghitung luas (area)
	luas := sisi * sisi

	// Menampilkan hasil
	fmt.Printf("Keliling Alun-alun Purwokerto: %.2f meter\n", keliling)
	fmt.Printf("Luas Alun-alun Purwokerto: %.2f meter persegi\n", luas)
}
