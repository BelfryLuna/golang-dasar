package main

import "fmt"

// Function, sebuah blok kode yang dibuat dalam program agar dapat dapat dipanggil berulang kali tanpa harus menuliskan blok kodenya lagi.
// keyword: func

// membuat function,
/* Berada di luar function main,
pembuatannya,
func nama_function() {
	...baris kode
}
*/
func sayHello()  {
	fmt.Println("Hello World!")
}


func passGrade (ujian int, absensi int) {
	if ujian > 80 {
		lulusUjian := true
		fmt.Println("Anda Lulus ujian")
	} else {
		lulusUjian := false
		fmt.Println("Anda tidak lulus ujian")
	}

	if absensi > 75 {
		lulusAbsensi := true
		fmt.Println("Anda lulus absensi")
	} else {
		lulusAbsensi := false
		fmt.Println("Anda tidak lulus absensi")
	}

	if lulu

}

func main() {

// memanggil function
/* tinggal ketikkan nama_function diikuti dgn kurung buka dan tutup
nama_function()
*/
sayHello()

}