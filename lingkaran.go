package main

import "fmt"

func main() {
	//variable
	var keliling, luas, r float64
	const pi = 3.14
	fmt.Println("input r")
	fmt.Scan(&r)
	luas = pi * r * r
	keliling = 2 * pi * r
	fmt.Println("Luas Lingkaran = ", luas)
	fmt.Println("Keliling Lingkaran = ", keliling)
}
