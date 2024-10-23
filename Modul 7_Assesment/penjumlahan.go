package main

import "fmt"

func main() {
	var x, y, j, jumlah int
	fmt.Scan(&x, &y)
	jumlah = 0
	for j = x; j <= y; j += 1 {
		jumlah = jumlah + j
	}
	fmt.Println("Hasil :", jumlah)
}
