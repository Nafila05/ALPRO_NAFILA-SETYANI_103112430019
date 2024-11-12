package main

import (
	"fmt"
	"strings"
)

func main() {
	var umur int
	var kewarganegaraan string

	fmt.Print("Masukkan umur Anda: ")
	fmt.Scan(&umur)

	fmt.Print("Masukkan kewarganegaraan Anda (contoh: WNI): ")
	fmt.Scan(&kewarganegaraan)
	kewarganegaraan = strings.ToUpper(strings.TrimSpace(kewarganegaraan))

	if umur >= 17 && kewarganegaraan == "WNI" {
		fmt.Println("Anda bisa mengikuti pemilu")
	} else {
		fmt.Print("Anda tidak bisa mengikuti pemilu karena ")
		if umur < 17 {
			fmt.Print("umur Anda belum mencapai 17 tahun")
		}
		if umur < 17 && kewarganegaraan != "WNI" {
			fmt.Print(" dan ")
		}
		if kewarganegaraan != "WNI" {
			fmt.Print("Anda bukan WNI")
		}
		fmt.Println(".")
	}
}
