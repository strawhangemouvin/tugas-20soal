package main

import "fmt"

func main() {
	// no 6
	hobi := []string{"Membaca", "Menulis"}
	hobi = append(hobi, "Berenang")
	fmt.Println(hobi)
	fmt.Println("Jumlah hobi:", len(hobi))

	// no 7
	// tebak output
	// [20 30 40]
	// [10 20]
	// [40 50]

}
