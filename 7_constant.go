package main

// CONSTANT
// adalah jenis variabel yang nilainya tidak bisa diubah. Inisialisasi nilai hanya dilakukan sekali di awal, setelah itu variabel tidak bisa diubah nilainya

import "fmt"

func main() {

	//  berbeda dengan variable biasa, saat deklarasinya menggunakan keyword const. disini saat mendeklaraskan sebuah const nilainya harus langsung di inisialisasi. jika tidak akan menghasilkan error

	// penulisan pertama, dengan lengkap
	const nama string = "Sonny"

	// kedua, lebih singkat tanpa tipe data krn const juga menggunakan konsep typr inference
	const namaDepan = "Sonny"

	
	// *uniknya variabel const tidak akan error jika tidak digunakan seperti variabel biasa


	// penulisan multiple const, mendeklarasikan const secara sekaligus
	const (
		namaBelakang = "Syahputra"
		namaDapur = "Ica"
	)

	fmt.Println(namaBelakang)
	fmt.Println(namaDapur)
}