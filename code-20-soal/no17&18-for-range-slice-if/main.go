package main

import "fmt"

func main() {
	// no 17
	nilai := []int{80, 65, 90, 70, 55}
	jmlRemedial := 0
	for _, n := range nilai {
		status := "Lulus"
		if n < 70 {
			status = "Remedial"
			jmlRemedial++
		}
		fmt.Printf("%d %s\n", n, status)
	}
	fmt.Printf("Jumlah remedial: %d\n", jmlRemedial)

	// no 18
	hewan := []string{"kucing", "ayam", "ikan", "ular", "ayam"}
	totalAyam := 0
	for i, h := range hewan {
		if h == "ayam" {
			fmt.Printf("ayam ditemukan di indeks %d\n", i)
			totalAyam++
		}
	}
	fmt.Printf("Total: %d\n", totalAyam)
}
