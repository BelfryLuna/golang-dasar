package database


// PACKAGE INITIALIZATION

// jadi package initialization itu merupakan fitur dimana kita dapat mengakses sebuah function secara otomatis (tanpa memanggilnya terlebih dahulu) ketika mengakses/mengimport package (packages tempat function itu dibuat)
// untuk mengakses fitur tsb menggunakan package initialization, atau membuat function dgn nama init() (private function)
// fitur ini berguna ketika, membuat function2 yang bertugas koneksi atau berkomunikasi dgn database2

var connection string

// ketika membuat function init, function ini otomatis akan dieksekusi ketika kita mengakses/mengimport package database ini di package lain
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
