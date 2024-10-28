package main

import (
	"fmt"
)

func main() {
	var angka int
	fmt.Print("Inputkan sebuah angka: ")
	fmt.Scan(&angka)

	if angka == 42 {
		fmt.Println("Hello, Universe")
	} else {
		fmt.Println(angka)
	}
}