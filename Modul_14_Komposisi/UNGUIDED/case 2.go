package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&n)

	factors := []int{}
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
			sum += i
		}
	}

	fmt.Printf("Faktor dari %d adalah: ", n)
	for i, factor := range factors {
		fmt.Print(factor)
		if i < len(factors)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()

	if sum == n {
		fmt.Printf("Ya, karena %d = ", n)
		for i, factor := range factors {
			fmt.Print(factor)
			if i < len(factors)-1 {
				fmt.Print(" + ")
			}
		}
		fmt.Printf(" = %d\n", n)
	} else {
		fmt.Printf("Tidak, %d bukan bilangan sempurna.\n", n)
	}
}
