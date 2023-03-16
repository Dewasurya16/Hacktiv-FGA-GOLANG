package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "selamat malam"
	words := strings.Fields(sentence)
	count := make(map[string]int)

	for _, word := range words {
		for _, char := range word {
			fmt.Printf("%c\n", char)
		}
		fmt.Println()
		count[word]++
	}

	fmt.Println(count)
}
