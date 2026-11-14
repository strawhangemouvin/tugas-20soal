package main

import "fmt"

func main() {
	// no 13
	kontak := map[string]string{
		"Rina": "0812",
		"Budi": "0813",
	}
	cari := "Sari"
	if nomor, ada := kontak[cari]; ada {
		fmt.Printf("Nomor %s: %s\n", cari, nomor)
	} else {
		fmt.Printf("Kontak %s tidak ditemukan\n", cari)
	}
	// Hasil jika 'cari' diganti "Budi": Nomor Budi: 0813


	// no 14
	// gula terdaftar, stok: 0
	// kopi tidak terdaftar
}
