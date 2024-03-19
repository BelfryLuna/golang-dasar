package main

import "fmt"

// Function dengan return value
/*
	ketika ingin membuat sebuah return value di dalam function, maka kita juga menetapkan tipe data dari nilai yang akan di return.

	ini wajib dibuat ketika kita mendeklrasikan tipe data nilai yang akan dikembalikan di sebuah function, maka kita wajib memberikan sebuah return value dgn tipe data yang sesuai

	tipe data (return value) dideklarasikan sesudah kurung() saat mebuat function.

	func nama_function (parameter) tipeData_dariReturnValue {

		return data_yangDikembalikan
	}
*/

func getName (name string) string {
	
	return name
}

func main() {

	// untuk menampilkan nya kita dapat menyimpannya dalam variabel lalu memanggilnya dalam Println()
	value := getName("Sonny")

	fmt.Println(value)

	//  atau menyimpannya langsung dalam Println()
	fmt.Println(getName("Eko"))
}