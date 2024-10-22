package main

// nama : christ daniel santoso
// nim: 2311102305
// SOAL NO 1
import (
	"fmt"
)

func hitungBiayaSewa(jam, menit int, isMember bool, voucher string) float64 {

	var tarifPerJam float64
	if isMember {
		tarifPerJam = 3500
	} else {
		tarifPerJam = 5000
	}

	durasiJam := float64(jam) + float64(menit)/60
	if durasiJam < 1 && menit >= 10 {
		durasiJam = 1
	}

	biayaDasar := durasiJam * tarifPerJam

	if len(voucher) >= 5 {
		biayaDasar *= 0.9
	}

	return biayaDasar
}

func main() {
	var jam, menit int
	var isMember bool
	var voucher string

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&isMember)
	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&voucher)

	biayaSewa := hitungBiayaSewa(jam, menit, isMember, voucher)

	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biayaSewa)
}
