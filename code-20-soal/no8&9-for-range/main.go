package main

import "fmt"

func main() {
	// no 8
	penjualan := [5]int{120, 150, 90, 200, 175}
	for i, p := range penjualan {
		fmt.Printf("Hari ke-%d: %d\n", i+1, p)
	}

	// no 9
	angka := [4]int{5, 10, 15, 20}
	total := 0
	for _, n := range angka {
		total += n
	}
	fmt.Println("Total:", total)

}
