package main

import "fmt"

func main() {
	// no 16
	hari := []string{"Senin", "Selasa", "Rabu"}
	idx := 5
	if idx >= 0 && idx < len(hari) {
		fmt.Println("Hari:", hari[idx])
	} else {
		fmt.Printf("Indeks %d di luar jangkauan, panjang slice: %d\n", idx, len(hari))
	}
	// Jika idx diganti 1, outputnya: Hari: Selasa

}
