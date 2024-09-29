package main

import "fmt"

// Fungsi F(x, y)
func F(x, y float64) float64 {
	return 1/(3*x*x+10) + 10*y + 7
}

func main() {
	// Mendefinisikan variabel x dan y
	var x, y float64

	// Input nilai x dan y dari user
	fmt.Println("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Println("Masukkan nilai y: ")
	fmt.Scan(&y)

	// Memanggil fungsi F dan mencetak hasilnya
	hasil := F(x, y)
	fmt.Printf("Hasil F(%.2f, %.2f) = %.4f\n", x, y, hasil)
}
