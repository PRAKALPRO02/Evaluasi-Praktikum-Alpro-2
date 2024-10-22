package main

// nama : christ daniel santoso
// nim : 2311102305
// SOAL NO 2

import "fmt"

func isPerfectNumber(num int) bool {
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i*i != num {
				sum += num / i
			}
		}
	}
	return sum == num
}

func findPerfectNumbers(start, end int) {
	fmt.Printf("Perfect numbers antara %d dan %d: ", start, end)
	for i := start; i <= end; i++ {
		if isPerfectNumber(i) {
			fmt.Printf("%d ", i)
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
