package main

import "fmt"

func main() {
	var panjang, lebar int
	var keliling, luas int

	fmt.Print("Masukkan panjang: ")
	fmt.Scan(&panjang)

	fmt.Print("Masukkan lebar: ")
	fmt.Scan(&lebar)

	// Hitung luas dan keliling
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	// Tampilkan hasil
	fmt.Printf("Luas persegi panjang: %d\n", luas)
	fmt.Printf("Keliling persegi panjang: %d\n", keliling)
}
