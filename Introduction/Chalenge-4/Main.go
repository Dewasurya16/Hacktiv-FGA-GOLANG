package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	// Data teman-teman di kelas
	teman1 := Teman{"Ikshan", "Jl. Sudirman No. 123", "Developer", "Ingin belajar bahasa pemrograman baru"}
	teman2 := Teman{"Mamat", "Jl. Thamrin No. 456", "Designer", "Ingin meningkatkan skill programming"}
	teman3 := Teman{"Sambo", "Jl. Gatot Subroto No. 789", "Data Analyst", "Tertarik dengan bahasa pemrograman Golang"}

	// Menampilkan data teman berdasarkan argument pada terminal
	args := os.Args[1:]

	for _, arg := range args {
		switch arg {
		case "1":
			displayTeman(teman1)
		case "2":
			displayTeman(teman2)
		case "3":
			displayTeman(teman3)
		default:
			fmt.Println("Tidak ada data teman dengan nomor absen", arg)
		}
	}
}

func displayTeman(t Teman) {
	fmt.Println("Nama:", t.Nama)
	fmt.Println("Alamat:", t.Alamat)
	fmt.Println("Pekerjaan:", t.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", t.Alasan)
	fmt.Println()
}
