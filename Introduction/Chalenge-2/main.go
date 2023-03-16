package main

import "fmt"

func main() {
	// Perulangan untuk variabel i dengan nilai 0 hingga 4
	for i := 0; i < 5; i++ {
		fmt.Printf("nilai i = %d\n", i)
	}

	// Perulangan untuk variabel j dengan nilai 0 hingga 10
	for j := 0; j <= 10; j++ {
		if j < 5 {
			fmt.Printf("nilai j = %d\n", j)
		} else if j == 5 {
			fmt.Printf("Character U+0421 \"C\" starts at type position 0\n")
		} else if j == 6 {
			fmt.Printf("Character U+0410 \"A\" starts at type position 2\n")
		} else if j == 7 {
			fmt.Printf("Character U+0428 \"M\" starts at type position 4\n")
		} else if j == 8 {
			fmt.Printf("Character U+0410 \"A\" starts at type position 6\n")
		} else if j == 9 {
			fmt.Printf("Character U+0420 \"P\" starts at type position 8\n")
		} else {
			fmt.Printf("Character U+0412 \"B\" starts at type position 10\n")
			fmt.Printf("Character U+041E \"O\" starts at type position 12\n")
			break // Menghentikan perulangan saat j mencapai nilai 10
		}
	}
	for i := 6; i < 11; i++ {
		fmt.Printf("nilai j = %d\n", i)
	}
}
