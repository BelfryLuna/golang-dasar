package main

import "fmt"

func main() {
	// Tipe data map, kurang lebih sama dengan array/slice (tipe data yang berisi kumpulan data dgn tipe data yang sama)
	// letak bedanya di tipe data map, kita dapat memodifikasi index (pada map adalah key) menjadi string, dll

	// dimana keywordnya,
	/* map [tipe_data_keynya] tipe_data_valuenya {
		"key" : "value",
		...
	} */

	person := map[string]string{
		"nama":   "Sonny",
		"alamat": "Jalan. Merdeka",
	}

	// atau... var person1 map[string]string = map[string]string {}

	fmt.Println(person)
	fmt.Println(person["nama"])

	// tidak menambahkan datanya secara langsung

	person3 := map[int]string {}

	// menambahkan datanya
	person3[1] = "Sonny"

	// mengubah datanya
	person3[1] = "Irsan"

	fmt.Println(person3[1])


	// build function di map

	// menghapus data di map dgn key
	delete(person, "nama")
	
	fmt.Println(person)

	person["name"] = "IRSABI"

	// melihat jumlah data per key
	fmt.Println(len(person))



}