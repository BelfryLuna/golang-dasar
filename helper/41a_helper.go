package helper

// Nah disini kita akan membuat sebauh package baru dgn nama helper,
/*
Langkahnya, 
	kita membuat folder lagi di dalam folder projek kita, lalu diberi nama(nama folder ini nantinya juga akan dijadikan nama dari package yang ingin kita buat) 
	lalu buat sebuah file (cth : helper.go), nah disini kita akan menuliskan nama package nya dgn nama yang sama dgn nama folder yang telah dibuat sebelumnya.
*/

// dalam package helper ini kita dapat membuat sebuah baris kode misalnya function, method dll sebebas mgkin. Yang nantinya baris kode yang dibuat dalam package helper ini dapat digunkan pada package lain (misalnya: main. krn dalam defaultnya di golang program yang ketika di build yang akan dijalankan atau di compile adalah baris kode dalam package main dan main function) yang nantinya akan diekseskusi disana.

func SayHello(name string) string {
	return "Hello " + name
}

func PassGrade(nilaiUjian int, absensi int) string{
	var hasilUjian, hasilAbsensi bool

	if nilaiUjian >= 80 {
		hasilUjian = true
	} else {
		hasilUjian = false
	}

	if absensi >= 75 {
		hasilAbsensi = true
	} else {
		hasilAbsensi = false
	}

	if hasilAbsensi && hasilUjian {
		return "Anda Lulus"
	} else {
		return "Anda tidak Lulus"
	}
}

/*dapat dilihat ketika kita membuat function di luar dari package main, misalnya dalam package helper ini nama dr function nya harus  di awali dgn huruf kapital*/

