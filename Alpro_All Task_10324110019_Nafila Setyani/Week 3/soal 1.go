package main

import "fmt"

func main() {
	var bilangan int
	var hasil bool

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	// Cek apakah bilangan ganjil menggunakan modulus (%)
	hasil = bilangan%2 != 0

	// Tampilkan hasil
	if hasil {
		fmt.Println(bilangan, "adalah bilangan ganjil")
	} else {
		fmt.Println(bilangan, "adalah bilangan genap")
	}
}
