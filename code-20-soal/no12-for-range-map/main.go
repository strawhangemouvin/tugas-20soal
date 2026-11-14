package main

import "fmt"

func main() {
	// no 12
	stok := map[string]int{
		"pensil": 10,
		"pulpen": 20,
		"buku":   5,
	}
	jenisBarang := 0
	totalStok := 0
	for _, jumlah := range stok {
		jenisBarang++
		totalStok += jumlah
	}
	fmt.Println("Jenis barang:", jenisBarang)
	fmt.Println("Total stok:", totalStok)


}
