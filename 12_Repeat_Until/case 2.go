package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("Masukkan kata: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(strings.ToLower(input))

		if input == "telkom" {
			fmt.Println("Program selesai.")
			break
		}
	}
}
