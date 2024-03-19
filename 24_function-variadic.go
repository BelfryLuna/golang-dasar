package main

import "fmt"

// variadic function
// kita dapat menjadikan salah satu parameter dr sebuah functiob menjadi sebuah variabel arguments
// parameter tersebut harus menjadi parameter satu2nya atau multiple parameter maka  parameter tersebut harus berada di akhir untuk dijadikan variabel arguments
// parameter yang menjadi var arguments dapat menampung lebih dari satu data dan selayaknya slice atau array sehingga dapat diiterasi
// Perbedaan antara mengirim data array atau variabel argumen
/*
	pada array, kita harus mendeklarasikan arraynya terlebih dahulu, dan mengirimkan variabel array tersebut ke dalam parameter
	sedangkan pada variabel argumen kita dapat langsung mengrimkan datanya tanpa mebuat array atau slice
	misalnya, getFunction (10,10,10) dengan koma jika datanya lebih dari sAtu
*/

// Untuk membuat varaibel arguments,
/*
	saat membuat parameter di function
		nama_function (numbers ...int)
	kita memberi titik(...) pada saat menuliskan tipe data dari parameter

	
*/

func sumAll(numbers ...int) int {
	total := 0

	// disini parameter numbers (yang merupakan variabel arguments) akan dianggap slice integer []int
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	// fmt.Println(sumAll([]int {10,10,10})) // mengirim data slice
	fmt.Println(sumAll(10,12,13)) // mengirim data langsung dalam var arguments

	// jika kita sudah terlanjur membuat slice dan ingin mengirimkan data slice tersebut dalam var arguments maka tambahkan ... pada variabel yang menyimpan slice tersebut

	numbers := []int {10,10,10}

	// seperti ini, maka data slice tersebut akan dikonver menjadi variabel arguments
	fmt.Println(sumAll(numbers...))

	// jika tidak akan errors, krn variabel argumen tidak dapat menerima data slice tersebut tanpa menambahkan ...
	fmt.Println(sumAll(numbers))
}