package main

import "fmt"

func hitungBiayaSewa_2311102181(jam int, menit int, member bool, voucher int) float64 {
	
	var tarif int
	if member {
		tarif = 3500
	} else {
		tarif = 5000
	}

	
	if menit >= 10 {
		jam++
	}

	
	total := float64(jam * tarif)

	if jam > 3 && (voucher == 5 || voucher == 6) {
		total *= 0.9 
	}

	return total
}

func main() {
	var jam, menit, voucher int
	var member bool

	
	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam)

	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)

	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&member)

	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&voucher)

	biaya := hitungBiayaSewa_2311102181(jam, menit, member, voucher)

	
	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.0f\n", biaya)
}
