package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Mengatur seed untuk angka acak berdasarkan waktu
	rand.Seed(time.Now().UnixNano())

	// Memilih angka acak antara 1 hingga 100
	randomNumber := rand.Intn(100) + 1
	var guess int
	maxAttempts := 5

	fmt.Println("Selamat datang di permainan tebak angka!")
	fmt.Println("Saya telah memilih sebuah angka antara 1 hingga 100.")
	fmt.Printf("Anda memiliki %d kesempatan untuk menebaknya.\n", maxAttempts)

	// Memberikan pengguna hingga 5 kali kesempatan untuk menebak
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		fmt.Printf("\nPercobaan %d: Masukkan tebakan Anda: ", attempts)
		fmt.Scan(&guess)

		if guess < randomNumber {
			fmt.Println("Terlalu kecil!")
		} else if guess > randomNumber {
			fmt.Println("Terlalu besar!")
		} else {
			fmt.Println("Selamat! Anda menebak angka dengan benar!")
			break
		}

		// Jika pengguna tidak berhasil menebak dalam 5 percobaan
		if attempts == maxAttempts {
			fmt.Printf("\nAnda kehabisan kesempatan! Angka yang benar adalah %d.\n", randomNumber)
		}
	}
}
