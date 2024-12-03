package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var total int

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Masukkan harga barang (ketik 0 untuk selesai): ")
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)
		price, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Masukkan angka yang valid.")
			continue
		}

		if price == 0 {
			break
		}

		total += price
	}

	fmt.Printf("Total belanja Anda adalah: %d\n", total)
}
