package main

import "fmt"

func main() {
	// Input harga beli tiga barang
	var hargaBeli1, hargaBeli2, hargaBeli3 float64
	fmt.Println("Masukkan harga beli tiga barang:")
	fmt.Scan(&hargaBeli1, &hargaBeli2, &hargaBeli3)

	// Hitung harga jual dengan keuntungan 5%
	keuntungan := 0.05
	hargaJual1 := hargaBeli1 + (hargaBeli1 * keuntungan)
	hargaJual2 := hargaBeli2 + (hargaBeli2 * keuntungan)
	hargaJual3 := hargaBeli3 + (hargaBeli3 * keuntungan)

	// Output harga jual
	fmt.Printf("Harga jual barang 1: %.2f\n", hargaJual1)
	fmt.Printf("Harga jual barang 2: %.2f\n", hargaJual2)
	fmt.Printf("Harga jual barang 3: %.2f\n", hargaJual3)
}
