package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var alas, tinggi float64
		fmt.Scan(&alas, &tinggi)    // Membaca input sisi alas dan tinggi dari segitiga
		luas := (alas * tinggi) / 2 // Rumus menghitung luas segitiga
		fmt.Println(luas)           // Menampilkan hasil luas dari segitiga
	}
}
