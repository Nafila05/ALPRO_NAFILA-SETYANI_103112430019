package main

import (
	"fmt"
)

func main() {
	var BelanjaAwal, diskon, total int

	fmt.Print("Masukkan total belanja awal: ")
	fmt.Scan(&BelanjaAwal)

	fmt.Print("Masukkan diskon : ")
	fmt.Scan(&diskon)

	total = BelanjaAwal - (BelanjaAwal * (diskon / 100))

	fmt.Printf("Total harga :", total)
}
