package main

import "fmt"

func main() {
	var totalPrice float64
	var choice string

	fmt.Println("Aplikasi Kasir Sederhana")

	for {
		var itemPrice float64
		fmt.Print("Masukkan harga barang: ")
		fmt.Scanln(&itemPrice)

		totalPrice += itemPrice
		fmt.Printf("Total sementara: %.2f\n", totalPrice)

		fmt.Print("Tambah barang lagi? (ketik 'tidak' untuk selesai): ")
		fmt.Scanln(&choice)

		if choice == "tidak" {
			break
		}
	}

	fmt.Printf("\nTotal belanja: %.2f\n", totalPrice)
	fmt.Println("Terima kasih!")
}
