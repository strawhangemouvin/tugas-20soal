package main

import "fmt"

func main() {
	// no 1
	angka := 17
	if angka%2 == 0 {
		fmt.Println(angka, "adalah bilangan genap")
	} else {
		fmt.Println(angka, "adalah bilangan ganjil")
	}

    // no 2
	nilai := 78
	var grade string
	if nilai >= 90 {
		grade = "A"
	} else if nilai >= 80 {
		grade = "B"
	} else if nilai >= 70 {
		grade = "C"
	} else {
		grade = "D"
	}
	fmt.Println("Grade:", grade)

    // no 3 (tebak output)
    // Sisa pembagian: 3

}
