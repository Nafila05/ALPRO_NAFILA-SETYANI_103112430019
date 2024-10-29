package main

import (
	"fmt"
)

func main() {
	var nilai float64

	fmt.Print("Masukkan nilai ujian: ")
	fmt.Scan(&nilai)

	if nilai >= 70 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}
}
