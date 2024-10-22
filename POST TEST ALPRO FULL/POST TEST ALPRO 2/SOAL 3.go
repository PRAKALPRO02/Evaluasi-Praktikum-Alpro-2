package main

// nama : christ daniel santoso
// nim : 2311102305
// SOAL NO 3

import "fmt"

func hitungPertemuan(x, y int) int {
	jumlahPertemuan := 0
	for i := 1; i <= 365; i++ {
		if i%x == 0 && i%y != 0 {
			jumlahPertemuan++
		}
	}
	return jumlahPertemuan
}

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	hasil := hitungPertemuan(x, y)
	fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", hasil)
}
