package main

import "fmt"

// POINTER [META]

// TEORI
// defaultnya dari Golang, dalam semua data di variable itu passing by value (duplikasi value) bukan by reference
	// artinya, ketika mengirimkan data variable secara langsung ke dalam sebuah function, method atau variable lain.
	//  maka data dr variabel (yang dikirimkan) tsb akan di copy/duplikasi ke dalam function, method ataupun variable nya.
	// hal ini menyebabkan data di variabel awal dan data function, method atau variable lain yang menerimanya, terpisah.
	//  jadi ketika kita mengubah datanya dr function, method, variable. maka, data dri variabel awal tidak akan berubah. yang berubah hanya data duplikasi yang tersimpan di function, method, atau vatiabel nya.


// disini kita mengimplementasikannya konsep pass by value menggunakan struct
type Adress struct {
	Kota, Propinsi, Negara string
}


func main() {
	
	// pertama, address akan membuat sebuah data dr struct sebelumnya
	address := Adress {"Rokan Hilir", "Riau", "Indonesia"}

	address2 := address // disini kita mengirimkan data dr address menjadi data di address2
	// disini yang terjadi, adress akan menduplikasi data yang ia miliki ke address2. jadi address dua datanya tidak merujuk pada data di address, melainkan telah mendapat data baru hasil duplikasi dr address. 
	// sehingga ketika kita mengubah data field di address2 dua, itu tidak akan memengaruhi data field di address

	// sederhananya antara address dan address2 memiliki value yang terpisah atau tersendiri

		// contoh, disini mengubah data field kota di address2
		address2.Kota = "Pekanbaru"
	
	fmt.Println(address) // disini data field kota di address tidak berubah
	fmt.Println(address2) // yang berubah hanya data di address2 yg merupakan hasil copy atau duplikasi nya


	// Ketika kita tidak ingin menggunakan konsep pass by value dikarenakan, 
	// ketika kita terlalu banyak menduplikasi data, maka akan memakan banyak memori yang nantinya memperlambat performa aplikasi.

	// hal ini dapat diatasi ketika, kita menggunakan konsep by reference, jadi kita mengirimkan data variabel ke wadah(function, method, variabel dll) baru jadi tidak ada duplikasi data. tetapi ketika merubah data maka akan merujuk data memori yang sama dgn variabel yang dikirimkan
	// ketika merubah data di wadah baru, ini akan merubah data keduanya dari variabel awal dan wadah baru

	// by reference ini dapat digunakan dngn fitur POINTER
		// POINTER > mengubah pass by value menjadi by reference

		// untuk menggunakan fitur Pointer, kita perlu operator &
		// Operator &, digunakan untuk membuat variabel yang memiliki nilai pointer ke variabel lain. 
		// caranya, ketikkan operator & dan diikuti variabel yg dituju> &nama_variabelDituju

	// dengan menggunakan pointer
		addressPointer1 := Adress {"Dumai", "Riau", "Indonesia"}

		addressPointer2 := &addressPointer1 // maka addressPointer2 akan menjadi variabel dengan nilai yng pointer ke addressPointer1

		// maka tipe data addressPointer2 adalah *Adress (pointer atau reference ke Adress)
		// addressPointer1 adalah Adress
		addressPointer2.Propinsi = "Maldive"
		fmt.Println(addressPointer1) // datanya ikut berubah mengikuti addressPointer 2
		fmt.Println(addressPointer2) // data yang dirubah pada reference ke memori yg sama addressPointer1, makanya addressPointer1 ikut berubah


}