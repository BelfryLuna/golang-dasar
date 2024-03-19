package main

// ERROR
/*
Sebelumnya error telah dibahas pada defer, panic, recover. ketiga itu digunakan dalam menangani error yang dapat menghentikan aplikasi.
ketika casenya hanya error biasa, barulah digunakan fitur error ini digunakan.
*/

// ERROR, ini di golang merupakan sebuah interface yang berisi sebuah method Error(). Error() ini digunakan untuk menangkap pesan error dari fungsi2 dr package error yang telah dibuat (cth: errors.New("...."))

// Error ini tersedia dalam package tersendiri, yang khusus untuk fitur error
import (
	"errors" // berikut nama package error nya
	"fmt"
)

func pembagiNol (nilai int, pembagi int) (int, error) /*disini kita memngembalikan nilainya dalam unt, dan jika terjadi error, kembalikan error nya apa. Dan error ini adalah sebuah interface (maka dapat diisi dgn nil)*/ {
	if pembagi == 0 {
		return 0, errors.New("Pembaginya adalah Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := pembagiNol(100, 5)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error()) // err.Error(), Error() ini adalah method bawaan dr interface error, yang digunakan untuk mengakses pesan error yang ada di dalam function tersebut (yaitu errors.New())
	}
}