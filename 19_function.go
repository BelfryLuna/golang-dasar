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

func main() {

// memanggil function
/* tinggal ketikkan nama_function diikuti dgn kurung buka dan tutup
nama_function()
*/
sayHello()

}