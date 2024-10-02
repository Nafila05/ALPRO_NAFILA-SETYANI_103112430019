package main

import "fmt"

func main() {
	var idr int
	const kurs float64 = 15000

	fmt.Print("Masukkan jumlah uang dalam IDR: ")
	fmt.Scan(&idr)

	if idr > 0 {
		usd := float64(idr) / kurs
		fmt.Printf("Jumlah dalam USD: %.2f\n", usd)
	} else {
		fmt.Println("Jumlah IDR harus bilangan bulat positif.")
	}
}
