package main

// CONSTANT
// adalah jenis variabel yang nilainya tidak bisa diubah. Inisialisasi nilai hanya dilakukan sekali di awal, setelah itu variabel tidak bisa diubah nilainya

import "fmt"

func main() {

	// 1, berbeda dengan variable biasa, saat deklarasinya menggunakan keyword const
	const nama string = "Sonny"
	// 2
	const namaDepan = "Sonny"

	
	fmt.Println(namaDepan)

	// 3
	const (
		namaBelakang = "Syahputra"
		namaDapur = "Ica"
	)

	fmt.Println(namaBelakang)
	fmt.Println(namaDapur)
}