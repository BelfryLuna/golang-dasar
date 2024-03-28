package main

import "fmt"

func main() {
	// operasi aritmatika hanya dapat diberlakukan atau digunakan pada tipe data number
	fmt.Println(12 % 2)

	angka1 := 12

	angka2 := 12

	angka1 += angka2 + 10  // augmented assigment dan tambah

	fmt.Println(angka1)

	angka1++  // increment

	fmt.Println(angka1)

	angka1 = -angka2  // negatif

	fmt.Println(angka1)

	angka3 := 12

	angka3 %= 2  // modulo augmented assignment

	fmt.Println(angka3)
}