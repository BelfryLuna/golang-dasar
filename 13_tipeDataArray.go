package main

// TIPE DATA ARRAY
/*
tipe data array adalah tipe data yang dapat menyimpan banyak data dgn tipe data yang sama dalam satu variabel(variabel array). 

*/

import "fmt"

func main() {
// di golang saat membuat array, kita terlebih dahulu menentukan jumlah data yang dapat ditampung dalam satu array dan juga menentukan tipe data nya. sehingga ketika menginput data yang melebihi jumlah atau tipe data yang berbeda maka akan menghasilkan error. tiap tiap data yang diinisialisasi tersimpan dan terurut dalam indeks (dimulai dr 0)


	// 1, disini di cara pertama kita terlebih dahulu menentukan jumalah data yang ditampung [5] kemudian tipe data nya string, [5]string seperti ini.
	var myArray1 [5]string

	myArray1[0] = "M. Sonny" // disini kita dapat mengisi data-data dala indeks tersebut, misalnya kita ingin mengisi index 0 dari myArray1, dgn cara seperti ini
	myArray1[1] = "Irsan"

	fmt.Println(myArray1)

	// 2, jika ingin menginisialisasi datanya secara langsung (horizontal)
	var myArray2 = [4]int {1,2,3}

	fmt.Println(myArray2)

	// 3 jika ingin menginisialisasi datanya secara langsung (horizontal)
	var myArray3 = [2]string {
		"Sonny",
		"Irsan",
	}

	fmt.Println(myArray3)

	var testArray = [6]string {
		"Sonny",
		"Irsan",
		"syahputra",
		"Sonny",
		"Sonny",
		"Sonny",
	}
	
	fmt.Println(len(testArray))


	// 4 jika kita belum pasti ketika hendak menginisialisasi jumlah data yang hendak dutampung,maka gunakan titik tiga [...]string seperti ini, namun harus diikuti dgn pengisian nilainya secara langsung. dari jumlah nilai yang diisi tsb akan menjadi jumlah daya tampung dr si array
	var myArray = [...]string{ // disini tidak ditetapkan jumlahnya brp, di situasi misalnya kita blm yakin ingin meambahkan berapa data ke dalam array tsb
		"M. Sonny",
		"Irsan",
		"Syahputra", // disini kita menginput tiga data, maka julah data array yg dapat ditampung array tsb berjumlah 3 mengikuti data yang diinputkan
	}

	fmt.Println(myArray[0])



	// array multidimensi ???
	arrayMulti := [3] [3]int{{1,2,3} , {4,5,6}, {7,8,9}}
	fmt.Println("array indeks ke 1", arrayMulti[1])

	array12 := [6]   [3]int{{1,2,3} , {12,5,6}, {32,12,3}, {32,12,1}, {12,1,3}, {1,2,3}}


}