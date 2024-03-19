package main

import "fmt"

// Functon sbg parameter
// function juga bisa dianggap sbg tipe data ketika kita membuat sebuah function sbg parameter
// untuk menuliskan, kita cukup menuliskan nama parameter untuk menampung function tsb kmudian diikuti dgn struktur function yang ingin dibuat.
// ( parameterLain, namaParameterFunction func(string) string )

func simpanParameter(name string, filter func(string) string) {
	filteredWord := filter (name)

	fmt.Println("Hello", filteredWord)
} 

// jadi kita dapat menginput sebuah function dengan struktur yang sama sebagai value dari parameter untuk parameter dgn tipe data function tsb (apabila strukturnya sama (satu parameter string dan satu return string))

// berikut function yang akan dimasukkan

func filterWord (name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

// jika function yang dijadikan tipe data saat menuliskan parameter kala pembuatan functon sudah kepanjangan. untuk menghindari ribet, kita dapat menggunakan type declaration pada construct/struktur function tsb

type AKA func(string) string

func iniFunction(nama string, callback AKA) {
	callbackup := callback (nama)
	fmt.Println("Hello",  callbackup) 
}

func callbackLagi(nama string) string {
	if len(nama) <= 2 {
		return "Nama tidak Valid"
	} else {
		return nama
	}
}

func main() {
	// untuk memanggilnya cukup panggil function awal, dan tuliskan function yang ingin dikirimkan tanpa kurung buka tutup() seperti function as value dari variabel

	simpanParameter("anjing", filterWord)

	// atau jika menyimpannya function yang ingin dikirimkan tsb dalam sebuah variabel, maka kita dapat mengirimkan variabel tsb sbg parameter

	value := filterWord

	simpanParameter("Eko", value)


	iniFunction("Ek", callbackLagi)
}