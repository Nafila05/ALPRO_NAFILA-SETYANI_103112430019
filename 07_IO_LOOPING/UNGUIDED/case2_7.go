package main

import (
	"fmt"
)

func main() {
	// Meminta input dari pengguna
	var jumlahBarang int
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlahBarang)

	// Inisialisasi total poin
	totalPoin := 0

	// Proses perhitungan poin
	if jumlahBarang <= 5 {
		totalPoin = jumlahBarang * 10
	} else {
		totalPoin = 5*10 + (jumlahBarang-5)*15
	}

	// Output hasil perhitungan
	fmt.Println("Total poin yang didapatkan:", totalPoin, "poin")
}
