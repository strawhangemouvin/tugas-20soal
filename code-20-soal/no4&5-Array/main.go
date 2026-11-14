package main

import "fmt"

func main() {
	// no 4
	buah := [5]string{"Apel", "Jeruk", "Mangga", "Pisang", "Anggur"}
	fmt.Println(buah)
	fmt.Println("Pertama:", buah[0])
	fmt.Println("Terakhir:", buah[len(buah)-1])

	// no 5
	// Penyebab error: Indeks 3 berada di luar jangkauan 
	// (out of bounds). Array [3]int hanya memiliki indeks 0, 1, dan 2. 
	// Mengakses angka[3] menyebabkan program panic.   Perbaikan (mencetak elemen terakhir, yaitu indeks 2):
	angka := [3]int{10, 20, 30}
	fmt.Println(angka[len(angka)-1])

}
