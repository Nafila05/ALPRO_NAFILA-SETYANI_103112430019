package main

import "fmt"

func main() {
	var alas, tinggi int
	fmt.Print("Masukkan alas segitiga: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)

	if alas > 0 && tinggi > 0 {
		luas := 0.5 * float64(alas) * float64(tinggi)
		fmt.Printf("Luas segitiga: %.2f\n", luas)
	} else {
		fmt.Println("Alas dan tinggi harus bilangan bulat positif.")
	}
}
