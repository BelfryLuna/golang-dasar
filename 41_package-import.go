package main

// PACKAGE DAN IMPORT

// Package adalah senuah tempat untuk mengorganisir kode program yng kita buat di golang
/*Mari anggap package seperti folder, dimana kita dapat membagi kode program kita dalam package-package atau folder tertentu*/
// menggunakan package-package artinya kita dapat merapikan kode program kita, jika kita harus membuat semua kode program kita yg sama hasilnya akan kebanyakan dan memperbanyak baris kode dalam satu file. dngn membagi-bagi mode program dalam package2 maka tentu akan merapikan baris kode kita.

// untuk memggunakan/mengakses baris kode(function dll) yang berada di package lain, maka kita harus mengimport dulu package tersebut
import (
	"mengulang-golang/helper" // disini kita akan menuliskan terlebih dahulu direktori (nama projek) dan nama package yang ingin diimport
	"fmt" // dan dapat lebih dr satu package dapat diimport (baik package)
)

func main() {
	impor := helper.PassGrade(89, 80) // dan menjalankan nya, mudah ketikkan dulu nama packagenya dan diikuti baris kode(function) yang ingin dijalankan cth: PassGrade()
	impor2 := helper.SayHello("Sonny")

	fmt.Println(impor2, impor) // sama seperti ini, nama packagenya fmt dan fungsi yang ingin dijalankan adalah Println()


	// ACCES MODIFIER
	helper.Namae = "Eko" // disini kita dapat megakses serta merubah data variabel (dengan nama diawali huruf kapital dri package helper)

	fmt.Println(helper.Namae)

	// fmt.Println(helper.named) // namun tidak dapat mengakses variabel yang namanya di awali huruf kecil dari package helper (mengembalikan undefined ketika di run)


	// PACKAGE INITIALIZATION
	
}