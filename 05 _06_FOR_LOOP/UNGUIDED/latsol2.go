package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah kerucut: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var r, t float64
		fmt.Printf("Masukkan jari-jari dan tinggi kerucut ke-%d (pisahkan dengan spasi): ", i+1)
		fmt.Scan(&r, &t)

		// Rumus volume kerucut: (1/3) * π * r^2 * t
		volume := (1.0 / 3.0) * math.Pi * r * r * t
		fmt.Printf("Volume kerucut ke-%d: %.2f\n", i+1, volume)
	}
}
