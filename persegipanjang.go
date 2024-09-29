package main

import "fmt"

func main() {
	//variable
	var keliling, luas, panjang, lebar int
	fmt.Println("input panjang, lebar")
	fmt.Scan(&panjang, &lebar)
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	fmt.Println("Luas Persegi Panjang = ", luas)
	fmt.Println("Keliling Persegi Panjang = ", keliling)
}
