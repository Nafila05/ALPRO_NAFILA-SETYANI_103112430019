package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	secretNumber := 7

	for {
		fmt.Print("Tebak angka (1-10): ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		guess, err := strconv.Atoi(input[:len(input)-1])

		if err != nil {
			fmt.Println("Masukkan angka yang valid.")
			continue
		}

		if guess == secretNumber {
			fmt.Println("Selamat, tebakan Anda benar!")
			break
		} else {
			fmt.Println("Tebakan Anda salah, coba lagi.")
		}
	}
}
