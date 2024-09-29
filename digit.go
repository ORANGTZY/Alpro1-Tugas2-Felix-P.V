package main

import "fmt"

func main() {
	var x int

	// Input bilangan bulat positif
	fmt.Println("Masukkan bilangan bulat positif (<= 999): ")
	fmt.Scan(&x)

	// Deklarasi digit pertama, kedua, dan ketiga
	var d1, d2, d3 int

	// Memisahkan digit
	if x < 10 {
		d1 = 0
		d2 = 0
		d3 = x
	} else if x < 100 {
		d1 = 0
		d2 = x / 10
		d3 = x % 10
	} else {
		d1 = x / 100
		d2 = (x % 100) / 10
		d3 = x % 10
	}

	// Cetak hasil digit
	fmt.Printf("Digit pertama (ratusan): %d\n", d1)
	fmt.Printf("Digit kedua (puluhan): %d\n", d2)
	fmt.Printf("Digit ketiga (satuan): %d\n", d3)
}
