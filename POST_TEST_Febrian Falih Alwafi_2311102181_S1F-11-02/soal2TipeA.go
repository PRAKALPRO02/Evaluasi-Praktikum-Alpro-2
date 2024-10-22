package main

import "fmt"

func isPerfectNumber_2311102181(num int) bool {
        sum := 0
        for i := 1; i < num; i++ {
                if num%i == 0 {
                        sum += i
                }
        }
        return sum == num
}

func findPerfectNumbers(start, end int) {
        fmt.Println("Perfect numbers antara", start, "dan", end, ":")
        for num := start; num <= end; num++ {
                if isPerfectNumber_2311102181(num) {
                        fmt.Print(num, " ")
                }
        }
        fmt.Println()
}

func main() {
        var a, b int
        fmt.Print("Masukkan nilai a: ")
        fmt.Scan(&a)
        fmt.Print("Masukkan nilai b: ")
        fmt.Scan(&b)
        findPerfectNumbers(a, b)
}