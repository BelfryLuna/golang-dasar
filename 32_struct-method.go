package main

import "fmt"

type Pelanggan struct {
	Nama, Alamat string
	Umur         int
}

// STRUCT METHODE
// kita dapat membuat function dgn menyomot template data dari struct, function yang menggunakan dan menempel di struct disebut method

// untuk membuat sebuah method sama seperti function biasa sebelem menuliskan nama function nya, kita menuliskan parameter_dariStruct (parameter ini ibaratnya seperti variabel yg ingin mengakses data struct dari object struct ke dalam method) dan nama_struct yang ingin ditempel di dalam kurung.

// seperti ini, dan sama seperti function pada umumnya kita dapat memberikan parameter dan tipedata pengembalian jika ingin menambahkan return value,
func (pelanggan Pelanggan) /*parameter_dariStruct (parameter ini ibaratnya seperti variabel yg ingin mengakses data struct dari object struct ke dalam method) yaitu pelanggan(parameter) dgn tipe data dr struct Pelanggan*/ sayHello(name string) {
	fmt.Println("Halo,", name, "nama saya", pelanggan.Nama, "umur saya", pelanggan.Umur) // seperti di pelanggan.Nama, parameter pelanggan ini digunakan jika ingin mengakses field dalam struct yang telah diisi dalam object/variabel ke sebuah method. dan kasusnya disini ingin mengakses data di field Nama dan lainnya
}



func main() {
	// kita tak dapat memanggil sebuah method seperti function pada umumnya, tapi mengaksesnya ke dalam sebuah object yang mengisi data pada struct

	// disini
	Joko := Pelanggan {
		Nama:  "Joko",
		Umur: 21,
	}

	Joko.sayHello("Joni") // untuk aksesnya tuliskan dulu nama object (variabel yang menyimpan isian struct) titk(.) dan nama methodnya.
	// dgn begini method dapat juga mengakses data dri field yang telah diisi di object.

	Sonny := Pelanggan {Nama: "Sonny", Umur: 22}

	Sonny.sayHello("Cici")

}