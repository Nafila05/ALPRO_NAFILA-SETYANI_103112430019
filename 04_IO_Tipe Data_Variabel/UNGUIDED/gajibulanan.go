package main

import "fmt"

func main() {
	var jamKerjaPerMinggu, upahPerJam float64
	const jamNormalPerMinggu = 40
	const mingguPerBulan = 4

	// Meminta input dari pengguna
	fmt.Print("Masukkan jumlah jam kerja per minggu: ")
	fmt.Scanln(&jamKerjaPerMinggu)

	fmt.Print("Masukkan upah per jam: ")
	fmt.Scanln(&upahPerJam)

	var gajiMingguan float64

	// Menghitung gaji mingguan
	if jamKerjaPerMinggu > jamNormalPerMinggu {
		lembur := jamKerjaPerMinggu - jamNormalPerMinggu
		gajiMingguan = (jamNormalPerMinggu * upahPerJam) + (lembur * 1.5 * upahPerJam)
	} else {
		gajiMingguan = jamKerjaPerMinggu * upahPerJam
	}

	// Menghitung gaji bulanan
	gajiBulanan := gajiMingguan * mingguPerBulan

	// Menampilkan hasil gaji bulanan
	fmt.Printf("Total gaji bulanan: %.2f\n", gajiBulanan)
}
