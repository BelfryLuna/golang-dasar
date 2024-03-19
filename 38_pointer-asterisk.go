package main

import "fmt"

// POINTER DAN ASTERISK OPERATOR

type Address struct {
	Kota, Propinsi, Negara string
}

func main() {

	address1 := Address {"Rokan Hilir", "Riau", "Indonesia"}

	// variabel pointer
	address2 := &address1 
	address3 := &address1
	// disini kita menjadikan address2 dan address3 sbg variabel pointer

	// meskipun kita tau bahwa addres2 adalah variabel pointer yag mereference/pointer ke address1, begitu kita mengubah datanya, maka hanya variabel pointer tersebutlah yg berubah. data pada variabel yg dipointing atau yg dijadikan referencing (address1) ataupun variabel pointer lain yng mengacu/referencing ke data address1 tidak akan berubah

	// contoh, kita mengubah data pada address2
	address2 = &Address {"Pekanbaru", "Riau", "Indonesia"} // krn address2 itu ointer, maka untuk mengubah datanya haruslah berupa pointer juga (tambahkan & sebelum data ubahnya). Dan disini address2 mengacu pada address memori yang baru atau mengacu pada pointer yang berbeda dgn addres1


	// disini kita lihat bahwa address2 tidak lagi mengacu pada data address1
	fmt.Println(address2)
	// dan address1 dan addres3 tidak berubah datanya
	fmt.Println(address1)
	fmt.Println(address3)



	// asterisk operator (*) ini digunakan jika kita ingin semua variabel pointer yang mengacu pada address1 (data yang di referencing/pointing diacu) dan address1 nya sendiri ikut berubah datanya dengan data dari salah satu variabel pointer (mengacu address1 juga) yang diubah datanya.

	// untuk menggunakan asterisk operator (*), kita cukup tambahkan bintang(*) sebelum nama variabel pointer yang hendak diubah datanya
	
	*address3 = Address {"Dumai", "Riau", "Indonesia"} // dan disini kita tidak perlu menambahkan pointer lagi pada datanya, krn ketika menggunakan asterisk(*) artinya kita mengubah valuenya dr memory(pointer) yang direference.

	// krn yang diubah adalah value dr pointernya/memorinya/ data yng diacu bukan pointernya. maka semua data baik itu semua variabel pointer yng mengacu pada address1(data acuannya) ataupun address1 nya sendiri akan ikut berubah datanya.

	fmt.Println(address3)
	fmt.Println(address1) // data address1 ikut berubah mengikuti address3` `
	fmt.Println(address2) // disini address2 tidak ikut berubah krn pointer atau acuannya telah berbeda krn sebelumnya datanya diubah tidak menggunakan asterisk operator

	// ibaratnya, masih pada memory yang sama tapiiii value dr memorinya diubah. sehingga semua data yng mengacu pada memory tersebut akan ikut berubah datanya mengikuti emorinya
}