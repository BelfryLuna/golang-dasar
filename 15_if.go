package main

import "fmt"

func main() {
	nilaiUjian := 80
	absensi := 76

	var lulusUjian bool
	var lulusAbsensi bool

	// If Statement [META]
	// if dan else

	if nilaiUjian >= 75 {
		lulusUjian = true
		fmt.Println("Anda lulus ujian")
	} else {
		lulusUjian = false
		fmt.Println("Anda tidak lulus ujian")
	}

	if absensi > 75 {
		lulusAbsensi = true
		fmt.Println("Dan Anda lulus absensi")
	} else {
		lulusAbsensi = false
		fmt.Println("Dan Anda tidak lulus absensi")
	}

	if lulusAbsensi && lulusUjian {
		fmt.Println("Maka, Anda lulus")
	} else {
		fmt.Println("Maka, Anda tidak lulus")
	}

	// if, else, else if

	sonny := 3

	if sonny % 2 == 0 {
		fmt.Println("Genap")
	} else if sonny % 2 == 1 {
		fmt.Println("Ganji")
	} else {
		fmt.Println("Ini dia")
	}

	// if short statement, membuat sebuah statement sederhana sebelum menuliskan kondisi. misalnya membuat sebuah variabel
	// yang kemudian digunakan untuk pengondisian

	// tanpa short statement
	nama := "Sonny"
	length := len(nama)

	if length > 4 {
		fmt.Println("Lebih besar")
	} else {
		fmt.Println("Lebih Kecil")
	}

	// dengan short statement, dimana short statement dengan kondisi dibatasi dgn titik koma
	if lenght := len(nama); lenght >= 5 {
		fmt.Println("Benar")
	} else {
		fmt.Println("Salah")
	}
}