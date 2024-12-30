package main

import "fmt"

func main() {
	var n_0019 int
	fmt.Print("Masukkan ukuran pola (harus bilangan ganjil): ")
	fmt.Scan(&n_0019)

	// Memastikan n adalah bilangan ganjil
	if n_0019%2 == 0 {
		fmt.Println("Ukuran harus bilangan ganjil.")
		return
	}

	// fungsi untuk membuat pola X
	for i := 0; i < n_0019; i++ {
		for j := 0; j < n_0019; j++ {
			// fungsi untuk menentukan posisi yang membentuk pola X
			if j == i || j == n_0019-i-1 {
				fmt.Print(i + 1) // Menampilkan angka sesuai baris
			} else {
				fmt.Print(" ") // Menampilkan spasi di posisi selain diagonal
			}
		}
		fmt.Println() // Pindah ke baris berikutnya
	}
}

// MULAI
//   INPUT n // Masukkan ukuran pola (harus bilangan ganjil)

//   JIKA n MOD 2 = 0 MAKA
//     CETAK "Ukuran harus bilangan ganjil."
//     BERHENTI
//   AKHIR JIKA

//   UNTUK i dari 0 hingga n - 1 LAKUKAN
//     UNTUK j dari 0 hingga n - 1 LAKUKAN
//       JIKA j = i ATAU j = n - i - 1 MAKA
//         CETAK i + 1 TANPA GANTI BARIS // Tampilkan angka sesuai baris
//       SELAIN ITU
//         CETAK " " TANPA GANTI BARIS // Tampilkan spasi
//       AKHIR JIKA
//     AKHIR UNTUK
//     CETAK GANTI BARIS // Pindah ke baris berikutnya
//   AKHIR UNTUK
// SELESAI
