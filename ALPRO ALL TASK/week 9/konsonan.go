package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	input = strings.ToLower(input)

	vowels := "aeiou"

	if len(input) == 1 && input[0] >= 'a' && input[0] <= 'z' {
		if strings.Contains(vowels, input) {
			fmt.Println("bukan huruf konsonan")
		} else {
			fmt.Println("huruf konsonan")
		}
	} else {
		fmt.Println("bukan huruf konsonan")
	}
}

// MULAI
//   DEFINISIKAN input SEBAGAI STRING
//   AMBIL NILAI input DARI PENGGUNA

//   KONVERSI input MENJADI huruf kecil (lowercase)

//   TETAPKAN vowels SEBAGAI "aeiou"

//   JIKA panjang input ADALAH 1 DAN input ADALAH huruf (a-z), MAKA
//     JIKA input TERMASUK DI DALAM vowels, MAKA
//       TAMPILKAN "bukan huruf konsonan"
//     JIKA TIDAK
//       TAMPILKAN "huruf konsonan"
//     AKHIR JIKA
//   JIKA TIDAK
//     TAMPILKAN "bukan huruf konsonan"
//   AKHIR JIKA
// SELESAI
