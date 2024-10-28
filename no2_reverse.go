package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func balikKata(kata string) string {
	runeKata := []rune(kata)
	for awal, akhir := 0, len(runeKata)-1; awal < akhir; awal, akhir = awal+1, akhir-1 {
		runeKata[awal], runeKata[akhir] = runeKata[akhir], runeKata[awal]
	}
	return string(runeKata)
}

func main() {
	fmt.Print("Masukkan Nama: ")
	inputScanner := bufio.NewScanner(os.Stdin)
	inputScanner.Scan()
	inputNama := inputScanner.Text()

	pecahanNama := strings.Fields(inputNama)

	for _, kata := range pecahanNama {
		fmt.Println(balikKata(kata))
	}
}