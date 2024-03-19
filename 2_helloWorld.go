package main // sebuah aplikasi/ main function (func main()) golang harus terdapat di dalam package main

import "fmt" // import ini akan mengambil sebuah function dr package bawaan golang (fmt), jika ingin menjalankan fungsi2 tertentu, misalnya Println() adalah fungsi yang terdapat dalam package fmt yang berguna untuk input i/o atau menampilakn tulisan di terminal. 
//  Dan juga kita juga dapat mengimport package lain2 baik dr bawaan golang, yang telah kita buat sendiri atau milik orang lain

func main() { //tiap tiap file go, harus memiliki main function, function utama yang akan dijalankan ketika program dijalankan
	fmt.Println("Hello World!")

	fmt.Println("Halo")
}