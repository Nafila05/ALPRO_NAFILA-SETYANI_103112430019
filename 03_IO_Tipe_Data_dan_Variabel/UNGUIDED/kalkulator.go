package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	// Input dari pengguna
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&num1)

	fmt.Print("Masukkan operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&num2)

	// Operasi aritmetika berdasarkan operator
	switch operator {
	case "+":
		fmt.Printf("Hasil: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Hasil: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Hasil: %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Hasil: %.2f\n", num1/num2)
		} else {
			fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
		}
	default:
		fmt.Println("Operator tidak valid.")
	}
}
