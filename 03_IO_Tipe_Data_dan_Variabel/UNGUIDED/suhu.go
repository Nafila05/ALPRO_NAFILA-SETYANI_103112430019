package main

import "fmt"

func main() {
	var fahrenheit float64

	// Input suhu dalam Fahrenheit
	fmt.Print("Masukkan suhu dalam derajat Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	// Konversi dari Fahrenheit ke Kelvin
	kelvin := (fahrenheit-32)*5/9 + 273.15

	// Menampilkan hasil konversi
	fmt.Printf("Suhu dalam Kelvin: %.2f\n", kelvin)
}
