package main

import "fmt"

// Function yang dapat mengembalikan beberapa (lebih dari satu) return value
// misalnya kita ingin mengemablikan beberapa return value (2 misalnya), maka kita perlu menuliskan 2 tipe data pengembalian untuk 2 return tersebut

// untuk membuat beberapa tipe data pengembalian, dibungkus dalam kurung dan tiap tipe data pengembalian dibatasi dgn koma


func getFullName() (string, string) {
	return "Eko", "Khannedy"
}

func getAllFullName() (string, string, string) {
	return "Sonny", "Irsan", "Syahputra"
}


func main() {
	fisrtName, lastName := getFullName() // untuk pemanggilan functionnya diperlukan deklarasi variabel sejumlah dgn return value yang. Dan masing2 variabel akan menyimpan masing2 return value yang ada. berdasarkan urutannya variabel pertama akan mengambil return value yang pertama dan variabel kedua akan mengambil return value yang kedua, dst

	// Jika ingin mengambil salah satu dr return value, kita dapat menghiraukan return value dengan cara yang sama jika tidak ingin mengambil index saat for range. yaitu, menambahkan _ (underscore) sesuai dengan urutannya.

	// jika ingin mengambil return yang pertama
	tryFirstName, _ := getFullName()

	// jika ingin mengambil return yang kedua
	_, tryLastName := getFullName()

	fmt.Println(tryFirstName)
	fmt.Println(tryLastName)


	fmt.Println(fisrtName, lastName)



	trySonny, tryIrsan, trySyahputra := getAllFullName()

	try1, _, _ := getAllFullName()

	fmt.Println(try1)
	fmt.Println(trySonny, tryIrsan, trySyahputra)
}