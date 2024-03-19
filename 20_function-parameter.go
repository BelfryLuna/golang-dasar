package main

import "fmt"

// Function dengan parameter
/*
	kita dapat menambahkan parameter ketika membuat sebuah function (optional), dan dapat lebih dari satu
	ketika kita membuat function yang memiliki parameter, saat memanggilnya kita wajib mengirimkan nilai sesuai dengan jenis data parameternya dan jumlah parameter yang dibuat 
*/

func sayHelloTo(namaDepan string, lastName string) {
	fmt.Println("Hello" , namaDepan, lastName)
}

func main() {
	sayHelloTo("Sonny", "Irsan")

	sayHelloTo("Eko", "Budi Nugraha")
}

