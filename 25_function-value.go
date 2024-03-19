package main

import "fmt"

// function sbg value
// jadi function di golang hampir sama seperti di javascript, kita dapat menyimpan function di dalam sebuah variabel.

// di JS, let simpan = iniFunction
// tanpa menambahkan () atau tidak memanggil function tsb
// dengan menyimpan function sbg value dari variabel, kita dapat memanggil variabel tsb layaknya function. intinya kita dapat menjalankan function yang ada di variabel tsb, dgn memanggil variabelnya

/*
	variabel_simpan := iniFunction

	fmt.Println(variabel_simpan()) // dipanggil layaknya function
*/

func getValue() string {
	return "Eko"
}

func getAllParameter(nama string) string {
	return nama
}

func main() {
// function as value from variable without parameter
  value := getValue

  fmt.Println(value())

//   function as value from variable with parameter
  simpan := getAllParameter

  fmt.Println(simpan("Sonny"))
}