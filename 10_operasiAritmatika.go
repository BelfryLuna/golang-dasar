package main

import "fmt"

func main() {

	
	fmt.Println(12 % 2)

	angka1 := 12

	angka2 := 12

	angka1 += angka2 + 10

	fmt.Println(angka1)

	angka1++

	fmt.Println(angka1)

	angka1 = -angka2

	fmt.Println(angka1)

	angka3 := 12

	angka3 %= 2

	fmt.Println(angka3)
}