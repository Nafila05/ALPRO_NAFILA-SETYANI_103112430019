package main

import "fmt"

func main() {

	var nilai int
	fmt.Print("Masukkan angka referensi: ")
	fmt.Scan(&nilai)

	var hitungNilai, hitungNol int

	fmt.Println("Masukkan 9 angka (0 atau nilai referensi):")
	for iterasi := 0; iterasi < 9; iterasi++ {
		var angka int
		fmt.Scan(&angka)

		if angka == nilai {
			hitungNilai++
		} else if angka == 0 {
			hitungNol++
		}
	}

	if hitungNilai > hitungNol {
		fmt.Println("Modus adalah: ", nilai)
	} else if hitungNilai < hitungNol {
		fmt.Println("Modus adalah: 0")
	} else {
		fmt.Println("Tidak ada modus dominan, kedua nilai muncul dengan jumlah yang sama.")
	}
}


Mulai
    // Deklarasikan variabel nilai, hitungNilai, dan hitungNol sebagai integer
    Deklarasikan nilai, hitungNilai, dan hitungNol sebagai integer
    
    // Tampilkan prompt untuk memasukkan angka referensi
    Tampilkan "Masukkan angka referensi: "
    
    // Baca nilai referensi dari input pengguna
    Baca nilai

    // Tampilkan prompt untuk memasukkan 9 angka
    Tampilkan "Masukkan 9 angka (0 atau nilai referensi):"
    
    // Iterasi untuk membaca 9 angka
    Untuk iterasi 0 sampai 8
        // Baca angka yang dimasukkan pengguna
        Baca angka
        
        // Jika angka sama dengan nilai referensi, increment hitungNilai
        Jika angka = nilai, hitungNilai++
        
        // Jika angka sama dengan 0, increment hitungNol
        Jika angka = 0, hitungNol++
    Selesai Untuk

    // Jika hitungNilai lebih besar dari hitungNol, modus adalah nilai referensi
    Jika hitungNilai > hitungNol
        Tampilkan "Modus adalah: ", nilai
    
    // Jika hitungNilai lebih kecil dari hitungNol, modus adalah 0
    Jika hitungNilai < hitungNol
        Tampilkan "Modus adalah: 0"
    
    // Jika tidak ada yang lebih besar, berarti tidak ada modus dominan
    Jika tidak
        Tampilkan "Tidak ada modus dominan."
Selesai
