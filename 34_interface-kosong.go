package main

import "fmt"

// GOLANG bukan bahasa yang menerapkan OOP

// pada golang, hampir sebagian besar function bawaan golang seperti :
/*
	Println(), panic(), recover() dll yang dapat menerima data jenis apapun, menjadikan interface kosong atau any sbg tipe data untuk  menampung segala jenis tipe data tersebut
*/
// Interface kosong, merupakan sebuah interface yang kita tidak mendefinisikan kontrak method apapun di dalamnya. sehingga interface tersebut bisa diisi dgn jenis data apapun.

// di golang yang baru, interface koong (interface{}) di aliaskan sbg any yang merupakan type declaration untuk interface kosong.

// jadi interface kosong atau any dapat kita jadikan sebuah tipe data (yang dapat menerima data apapun) untuk tipe data pengemabalian return value dari function, tipe data parameter dari sebuah function atau tipe data untuk variabel dll

// interface kosong atau any pada function

func ups() interface{} {
	// return 1
	// return "Halo"
	return true
}

func testAny(spesial any)  {
	fmt.Println(spesial)
}


func main() {
	// pada variabel

	var Sonny interface{}

	Sonny = "Sonny"
	Sonny = 12

	fmt.Println(Sonny)


	var cici any
		cici = "Sonny"
		cici = 13

	fmt.Println(cici)

	testAny(123)


	newSlice := make([]any, 2, 5)
	newSlice[0] = 12
	newSlice[1] = "Sonny"

		fmt.Println(newSlice)

	arrayNew := [...]any {
		"Sonny",
		1,
		"true",
	}

	fmt.Println(arrayNew)
}