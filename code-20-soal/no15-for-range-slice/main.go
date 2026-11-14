package main

import "fmt"

func main() {
	// no 15
	nilai := []int{80, 65, 90, 70}
	total := 0
	for _, n := range nilai {
		total += n
	}
	rataRata := float64(total) / float64(len(nilai))
	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Rata-rata: %.2f\n", rataRata)
}
