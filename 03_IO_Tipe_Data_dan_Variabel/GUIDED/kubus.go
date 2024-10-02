package main

import "fmt"

func main() {
	var sisi int
	fmt.Print("Masukkan panjang sisi kubus: ")
	fmt.Scan(&sisi)

	if sisi > 0 {
		fmt.Printf("Volume kubus: %d\n", sisi*sisi*sisi)
	} else {
		fmt.Println("Panjang sisi harus positif.")
	}
}
