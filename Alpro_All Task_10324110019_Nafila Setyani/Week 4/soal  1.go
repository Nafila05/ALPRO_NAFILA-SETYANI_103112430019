package main

import "fmt"

type Transaksi struct {
	NamaBarang  string
	Jumlah      int
	HargaSatuan float64
}

func main() {
	var transaksi Transaksi
	var totalHarga float64

	fmt.Print("Nama Barang: ")
	fmt.Scan(&transaksi.NamaBarang)

	fmt.Print("Jumlah: ")
	_, err := fmt.Scan(&transaksi.Jumlah)
	if err != nil {
		fmt.Println("Input jumlah tidak valid")
		return
	}

	fmt.Print("Harga Satuan: Rp ")
	_, err = fmt.Scan(&transaksi.HargaSatuan)
	if err != nil {
		fmt.Println("Input harga satuan tidak valid")
		return
	}

	// Hitung total harga
	totalHarga = float64(transaksi.Jumlah) * transaksi.HargaSatuan

	// Tampilkan hasil
	fmt.Printf("Total Harga: Rp %.2f\n", totalHarga)
}
