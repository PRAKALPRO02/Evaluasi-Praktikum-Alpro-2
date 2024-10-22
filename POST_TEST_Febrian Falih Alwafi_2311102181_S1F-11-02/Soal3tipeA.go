package main

import "fmt"

func hitungPertemuanRahasia_2311102181(x, y, totalHari int) int {
       
        jumlahPertemuan := 0
        for hari := 1; hari <= totalHari; hari++ {
               
                if hari%x == 0 && hari%y != 0 {
                        jumlahPertemuan++
                }
        }
        return jumlahPertemuan
}

func main() {
        var x, y int
        const totalHari = 365 
        fmt.Print("Masukkan nilai x: ")
        fmt.Scan(&x)
        fmt.Print("Masukkan nilai y: ")
        fmt.Scan(&y)

        jumlahPertemuan := hitungPertemuanRahasia_2311102181(x, y, totalHari)
        fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", jumlahPertemuan)
}

