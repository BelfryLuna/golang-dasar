package main

// GOLANG (Go Language)
// Bahasa pemrograman yang dibuat oleh google dgn menggunakan bahasa C di tahun 2009
// Bahasa Golang populer krn:
/*
	dapat membuat teknologi-teknologi sperti: docker, kubernetes, dll
*/
// Golang umumnya digunakan untuk pembuatan API back-end

// Kenapa belajar bahasa Golang:
/*
	bahasanya yang sederhana
	mendukung concurrency programming
	mendukung garbage collector
	Sedang naik daun
*/


// Menginisialisasi project / membuat projek dengan Go Modules
/*
	command :  go mod init nama_projek
	*saat membuat projek, programmer golang memiliki kebiasaan menuliskan nama projeknya sama dengan nama direktori/folder tempat menyimpannya

	command ini akan menghasilkan file go.mod. File ini digunakan oleh Go toolchain untuk menandai bahwa folder di mana file tersebut berada adalah folder project. Jadi jangan dihapus ya file tersebut.

keyword: go toolchain
*/

// Command2 lain yang ada di go lang
/*
	// go run
	// go run nama-file.go 
	   Untuk menjalankan/mengeksekusi file dgn ekstensi .go yang berisi baris kode yang telah kita buat dan harus berada dalam package main	

	// go build
	   command untuk mengeksekusi sekaligus mengompile/kompilasi file .go menjadi binary file (.exe)
	   command go build, akan menghasilkan file baru pada folder yang sama, yaitu mengulang-golang.exe, yang kemudian dieksekusi. Default-nya nama project akan otomatis dijadikan nama binary.

	   atau kita bisa menentukan nama dari binary file hasil dr go build
	   command : go build -o nama_yang_diinginkan.exe	
*/

func main() {
	
}