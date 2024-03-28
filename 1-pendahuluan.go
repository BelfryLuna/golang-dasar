package main

// GOLANG (Go Language)
// Bahasa pemrograman yang dibuat oleh google dgn menggunakan bahasa C dan C++ di rilis tahun 2009
// Bahasa Golang populer krn:
/*
	dapat membuat teknologi-teknologi sperti: docker, kubernetes, dll
*/
// Golang umumnya digunakan untuk pembuatan API back-end

// Kenapa belajar bahasa Golang:
/*
	bahasanya yang sederhana
	mendukung *concurrency programming dengan pengaplikasian yang cukup mudah
	mendukung *garbage collector
	Sedang naik daun
	kompilasi yang cepat
*/


// Menginisialisasi project / membuat projek dengan Go Modules
/*
	Modules digunakan untuk menginisialisasi sebuah project, sekaligus melakukan manajemen terhadap 3rd party atau library lain yang dipergunakan.

	command :  go mod init nama_projek
	*saat membuat projek, programmer golang memiliki kebiasaan menuliskan nama projeknya sama dengan nama direktori/folder tempat menyimpannya. Dengan ini telah menginisialisasi direktori/folder tersebut sebagai sebuah project Go dengan nama yang telah dibuat saat menuliskan command go mod init 

	command ini akan menghasilkan file go.mod. File ini digunakan oleh Go toolchain untuk menandai bahwa folder di mana file tersebut berada adalah folder project. 

keyword: *go toolchain
*/





// Command2 lain yang ada di go lang
/*
	// go version
	melihat versi golang yang tellah terinstall
	
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