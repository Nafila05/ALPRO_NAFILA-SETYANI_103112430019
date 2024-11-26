package main

import (
	"fmt"
)

func main() {
	const correctPassword = "nafila123"
	const maxAttempts = 3

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		var input string
		fmt.Printf("Masukkan password (percobaan %d/%d): ", attempts, maxAttempts)
		fmt.Scanln(&input)

		if input == correctPassword {
			fmt.Println("Login berhasil!")
			return
		}
		fmt.Println("Password salah.")
	}

	fmt.Println("Login ditolak")
}
