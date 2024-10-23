package main

import "fmt"

func main() {
	var N, j int
	fmt.Scan(&N)
	for j = 1; j <= N; j += 1 {
		fmt.Print(j*j, " ")
	}
}
