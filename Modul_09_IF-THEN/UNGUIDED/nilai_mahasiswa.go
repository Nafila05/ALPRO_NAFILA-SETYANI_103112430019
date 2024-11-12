package main

import (
	"fmt"
)

func klasifikasiNilai(nilai int) string {
	if nilai > 90 {
		return "A"
	} else if nilai >= 80 && nilai <= 90 {
		return "AB"
	} else if nilai >= 70 && nilai < 80 {
		return "B"
	} else {
		return "C"
	}
}

func main() {
	var nilai int
	fmt.Print("Masukkan nilai mahasiswa: ")
	fmt.Scan(&nilai)

	indeks := klasifikasiNilai(nilai)
	fmt.Printf("Indeks nilai: %s\n", indeks)
}
