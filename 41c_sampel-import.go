package main

import (
	"mengulang-golang/database" // ketika kita mengimport package, secara otomatis function init di package yang diimport akan dieksekusi 

	// ketika kita mengimport sebuah package namun tidak mengakses function apapun, maka golang tidak mengizinkannya. mengembalikan error. ketika kita hanya ingin menjalankan init function di dala package tsb dan tidak ingin mengakses function yang lain. maka gunakan blank identifier (_) di awal saat mengimport package nya
	_"mengulang-golang/internal" // seperti ini
	
	"fmt"
)

func main() {
   fmt.Println("test package")
   fmt.Println(database.GetDatabase())	
}