package main

import "fmt"

func main() {
	// no 20
	barang := []string{"kopi", "teh", "kopi", "roti", "teh", "kopi"}
	jumlah := []int{2, 1, 3, 4, 2, 1}

	rekap := map[string]int{}
	for i, b := range barang {
		rekap[b] += jumlah[i]
	}
	fmt.Println(rekap)

	var terlaris string
	var max int = 0
	for nama, jml := range rekap {
		if jml > max {
			max = jml
			terlaris = nama
		}
	}
	fmt.Printf("Terlaris: %s (%d)\n", terlaris, max)
}
