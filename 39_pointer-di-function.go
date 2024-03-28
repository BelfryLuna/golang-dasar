package main

import "fmt"

type AlamatKTP struct {
	Kelurahan, Kecamatan, Kabupaten string
}

func changeAddress(alamat AlamatKTP) {
	alamat.Kelurahan = "Bangko Permata" // melakukan perubahan value pada salah satu field di struct
	fmt.Println(alamat)
}


func changeAddressPointer (alamat *AlamatKTP) { // [solution] disini kalau ingin datanya berubah layaknya menggunakan pointer, jadikan tipe datanya menjadi pointer dgn menambhkan bintang(*) sebelum tipe datanya.

	alamat.Kabupaten = "Bengkalis"
	fmt.Println(alamat)
}

func main() {
	alamat := AlamatKTP{"Bangko Jaya", "Bangko Pusako", "Rokan Hilir"} 

	changeAddress(alamat) // data dr alamat(global) akan didke dalam parameter yang ada di function (sehingga ketika kita melakukan perubahan data alamat di dalam function)

	fmt.Println("Ini data di luar function," , alamat) // hal ini tidak akan memengaruhi data pada alamat(global var), karena merujuk memori data yang berbeda

	fmt.Println()



	// [SOLUTION] dengan menambahkan konsep pointer ( & / * / new() ), maka kita data yang sebelumnya tak berubah akan dapat berubah

	var alamatKtp AlamatKTP = AlamatKTP{"Bangko Permata", "Bangko Pusako", "Rokan Hilir"} // jika ingin memasukkan data yang bukan pointer
		changeAddressPointer(&alamatKtp) // maka tambahkan operator &
		fmt.Println(alamatKtp)
	
	var alamatKedua *AlamatKTP = &AlamatKTP{"Batang Dui", "Mandau", "Rohania"} // jika sudah terlanjur membuat data pointer
		changeAddressPointer(alamatKedua) // tambahkan seperti biasa
		fmt.Println(alamatKedua)

	// CONTOH LAIN
	alamatBaruLagi := &AlamatKTP{} // pointer
		changeAddressPointer(alamatBaruLagi)

	alamatKtp2 := AlamatKTP {"Bangko Permata", "Bangko Pusako", "Rokan Hilir"} // bukan pointer
		changeAddressPointer(&alamatKtp2)
	
	alamatBaru := new(AlamatKTP) // pointer
		changeAddressPointer(alamatBaru)

}