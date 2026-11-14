package main

import "fmt"

func main() {

	// no 10
	harga := map[string]int{
		"Kopi": 15000,
		"Teh":  8000,
		"Roti": 12000,
	}
	fmt.Println("Harga Teh:", harga["Teh"])
	harga["Susu"] = 10000
	harga["Kopi"] = 18000
	fmt.Println("Harga Kopi:", harga["Kopi"])
	fmt.Println("Jumlah menu:", len(harga))

	// no 11 (tebak output)
	// 80
	// 0
	// 1

}
