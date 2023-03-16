package main

import "fmt"

func main() {
	// Menampilkan nilai i: 21
	i := 21
	fmt.Printf("%v\n", i)

	// Menampilkan tipe data dari variabel i
	fmt.Printf("%T\n", i)

	// Menampilkan tanda %
	fmt.Printf("%%\n")

	j := true
	fmt.Printf("%v\n", j)

	// Menampilkan unicode russia: Я (ya)
	fmt.Printf("%c\n", '\u042F')

	// Menampilkan nilai base 10: 21
	fmt.Printf("%d\n", i)

	// Menampilkan nilai base 8: 25
	fmt.Printf("%o\n", i)

	// Menampilkan nilai base 16: f
	fmt.Printf("%x\n", 15)

	// Menampilkan nilai base 16: F
	fmt.Printf("%X\n", 15)

	// Menampilkan nilai base 10 dari karakter unicode: 1071
	fmt.Printf("%d\n", '\u042F')

	// Menampilkan float: 123.456000
	k := 123.456
	fmt.Printf("%f\n", k)

	// Menampilkan float scientific: 1.234560E+02
	fmt.Printf("%e\n", k)
}
