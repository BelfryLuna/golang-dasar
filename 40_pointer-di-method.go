package main

import "fmt"

// POINTER DI METHOD [SANGAT META UNTUK DILAKUKAKAN]

// Wajib hukumnya ketika membuat sebuah method kita perlu menjadikannya dalam pointer, krn secara default meskipun mengakses struct yang sama ketika dalam method kita mengubah salah satu data field dalam structnya, maka yang akan diubah adalah data duplikasinya. 

type Man struct {
	isMarried bool
}

func (man *Man) Married() { // caranya tambahkan (*) atau jadikan tipe datanya dr struct menjadi pointer dgn menambahkan (*)
	man.isMarried = false
}

func main() {
	sonny := Man{true}
	sonny.Married()
	/*
	maka ketika sebelumnya telah men-assign data field nya true, ketika kita mengakses method dalam variabel/object dr struct yang sama dimana di dalamnya kita mengubah isi dr field yang tdinya sudah di assign. dngn menggunakan konsep pointer, maka data di field tadi juga ikut berubah
	*/

	fmt.Println(sonny.isMarried)
}