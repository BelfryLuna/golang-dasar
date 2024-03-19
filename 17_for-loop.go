package main

import "fmt"

func main() {
	// For digunakan untuk perulangan

	// #1 sederhana hanya dengan kondisi saja
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)

		counter++ 
	}

	// #2 for dengan statement
	//      a. init statement dan b. post statement

	// dengan init statement dan post statement
	// init statemen; kondisi; post statement
	/* pertama,
			eksekusi init statement (sekali di awal perulangan)
			cek kondisi (apakah bernilai true); jika true
			eksekusi baris kode di dalam block
			eksekusi post statement
		kedua,
			cek kembali kondisi, jika true
			eksekusi baris kode di dalam block
			eksekusi post statement
		ketiga,
			ulangi hingga kondisi bernilai false
	*/
	for initStatement := 10; initStatement < 15; initStatement++ {
		fmt.Println("Perulangan ke ", initStatement)
	}

	// melakukan perulangan pada objek iterasi (array, slice, map)

	// 1. manual

	myArray := [5]string {"Sonny", "Erika", "Cici", "Ica", "Adinda"}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	// 2. menggunakan for range
	// index : menampung index/keys
	// name : menampung data pada setiap iterasi

	// cth pada array
	for index, name := range myArray {
		fmt.Println("Index ", index, "=", name)
	}

	// pada map

	newMap := map[string]string {
		"Pemain1" : "Sonny",
		"Pemain2" : "Cici",
	}

	for keys, nama := range newMap {
		fmt.Println("Keys ", keys, "= ", nama)
	}

	// pada slice

	newSlice := []int {
		1,
		3,
		4,
	}

	for index, nomor := range newSlice {
		fmt.Println("Index", index, "=", nomor)
	}

	// jika tidak butuh index atau keys nya (diganti jadi _ (undrscore))
	for _, nomor := range newSlice {
		fmt.Println("Ini, " , nomor)
	}

}