package main

import "fmt"

// STRUCT
// struct adalah sebuah template data (nantinya template data ini akan diisi dgn data) dimana dalam template tersebut kita dapat menggabungkan dan menyimpan beberapa jenis data dngn tipe data yang berbeda2 
// struct ini adalah representasi data yang nantinya disimpan dalam program aplikasi yang kita buat.
// nah data dgn jenis tipe data yang beragam tadi disimpan dalam field-field yang sebelumnya dibuat dan juga ditentukan tipe datanya masing-masing. jadi kita dapat menyimpan data dan disesuaikan dngn tipe data dari field yg akan kita tujukan sbg tempat menyimpan data tsb
// Jadi struct ini adalah sekumpulan field
// field ibaratnya adalah atribut 


// untuk membuat sebuah struct, kita menggunakan type declaration, diikuti nama untuk struct nya diikuti dgn keyword struct dan blok untuk mendeclare field field nya 
/* type nama_struct struct {
	nama_field tipeData_field
	dst...
}
*/
type Costumer struct {
	Nama, Address string
	Age           int
}

func main() {

	// dalam mengisi data template data struct tersebut, kita seolah dapat mengisi dgn cara disimpan dalam variabel yang dideklarasikan dgn tipe data dari nama stuct yg telah dibuat sebelumnya. dimana nanti struct ini tak ubahnya dgn tipe data lainnya seperti map, namun atribut/field di dalamnya telah ditetapkan sebelumnya saat pembuatan struct.

	// disini kita mengisi data struct costumer ke dalam variabel sonny
	// #1 ini adalah cara pertama, kita terlebih dahulu mendeklarasikan variabelnya. kemudian mengisi data tiap fieldnya. seperti map, array atau slice (namun caranya aga berbeda)
	var sonny Costumer // variabel akan dijadikan sebuah object (struct)
	sonny.Nama = "Sonny" // mengisi datanya tuliskan namavaribel.namaField
	sonny.Address = "Jalan. Riau"
	sonny.Age = 22

	fmt.Println(sonny) // mengakses struct seperti akses variabel biasa

	// print data struct per fieldnya (sama seperti jika ingin mengisi datanya)
	fmt.Println(sonny.Address)


	// membuat/mengisi data/object dari struct secara langsung saat pendeklarasian variabel structnya (literals). cth.. seperti pembuatan map secara langsung.

	// #2 menuliskan pasangan field dan value layaknya MAP
	var eko Costumer = Costumer{
		Nama: "Eko",
		Address: "Indonesia",
	}

	Jana := Costumer{
		Nama: "Jana",
		Age: 21,
	}
	
	fmt.Println(eko)
	fmt.Println(Jana)

	// #3 menuliskan value fieldnya secara langsung dan harus sesuai urutan, dan juga kita harus mengisi data fieldnya sesuai dgn jumlah field yang ada secara langsung. jika satu field tidak mendapat value maka akan error
	Janu := Costumer{"Sonny", "Pekanbaru", 23}

	fmt.Println(Janu)
}