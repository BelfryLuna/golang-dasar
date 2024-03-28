package main

import "fmt"

func main() {

	fmt.Println(len("Sonny")) // len, dapat digunakan untuk melihat jumlah karakter di dalam string
	fmt.Println("Sonny"[4]) // mengambil salah satu karakter dr string berdasarkan indexnya, dan mengembalikannya dalam biner

	var awal string // string di golang direpresentasikan dgn tanda petik dua ""
	// dan default value/zero value dr string adalah "" / string kosong. isalnya ketika kita tidak mengirimkan nilai pada variabel dgn tipe data string. maka variabel tersebut akan otomatis berisi zero/default value dr tipe data string tsb

	awal = "Sonny"

	fmt.Println(awal)
}