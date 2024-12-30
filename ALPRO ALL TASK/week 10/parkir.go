package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel
	var h1, m1, h2, m2 int

	// Input jam dan menit mulai parkir serta selesai parkir
	fmt.Print("Masukkan jam dan menit mulai parkir (h1 m1): ")
	fmt.Scan(&h1, &m1)
	fmt.Print("Masukkan jam dan menit selesai parkir (h2 m2): ")
	fmt.Scan(&h2, &m2)

	// Konversi waktu ke menit
	start := h1*60 + m1
	end := h2*60 + m2

	// Hitung durasi parkir dalam menit
	duration := end - start

	// Jika durasi negatif, artinya waktu parkir berada di luar jam operasional
	if duration < 0 {
		fmt.Println("Waktu parkir tidak valid!")
		return
	}

	// Konversi durasi ke jam dan menit
	hours := duration / 60
	minutes := duration % 60

	// Output hasil
	fmt.Printf("Durasi parkir: %d jam %d menit\n", hours, minutes)
}

//MULAI
//INPUT h1, m1, h2, m2

//start ← (h1 * 60) + m1
//end ← (h2 * 60) + m2

//duration ← end - start

//JIKA duration < 0 MAKA
//CETAK "Waktu parkir tidak valid!"
//BERHENTI
//AKHIR JIKA

//hours ← duration DIV 60
//minutes ← duration MOD 60

//CETAK "Durasi parkir:", hours, "jam", minutes, "menit"
//SELESAI//
