package main

import "fmt"

func main() {
	// no 19
	siswa := []map[string]string{
		{"nama": "Rina", "status": "aktif"},
		{"nama": "Budi"},
		{"nama": "Sari", "status": "nonaktif"},
	}
	for _, s := range siswa {
		status, ada := s["status"]
		if ada {
			fmt.Printf("%s: %s\n", s["nama"], status)
		} else {
			fmt.Printf("%s: status belum diisi\n", s["nama"])
		}
	}
}
