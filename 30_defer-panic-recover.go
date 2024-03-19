package main

import "fmt"

// CARA PENANGANAN ERROR YANG MENGHENTIKAN APLIKASI

// Defer Panic Recover
// DEFER
// merupakan sebuah fungsi yang digunakan dalam pemanggilan sebuah function. Dimana dgn menggunakan fungsi defer, kita dapat menjadwalkan eksekusi function yang menggunakan defer di akhir setelah semua function selesai dieksekusi, barulah function dgn defer tsb dieksekusi.
// jadi ini adalah sebuah fitur untuk memanggil sebuah function setelah sebuah function selesai.




func logging() {
	fmt.Println("Logging telah selesai")
}

func endPoint() {
	defer logging()  // defer akan dituliskan di awal baris kode sebuah function, namun tetap akan dieksekusi di akhir.

	// dan juga misalnya ada error di bawah baris kode defer tsb, maka function defer 
	fmt.Println("Aplikasi berhasil")
}

// PANIC
// panic adalah sebuah fungsi yang dapat ditempatkan dalam function yang berguna untuk menghentikan program ketika terjadi panic(error).
// ketika fungsi panic dipanggil maka program akan dihentikan. namun, defer function di awal akan tetap dijalankan.
// dalam membuat keyword panic kita dapat memberikan pesan (dapat dituliskan langsung atau melalui variabel)

func endApp() {
	fmt.Println("End APP")
}

func includePanic(error bool) {
	defer endApp()

	if error {
		panic("Error")
	}
}

// RECOVER
// sebuah fungsi untuk menangkap dan menyimpan data dari fungsi panic
// cara mengimplementasikannya, fungsi recover disimpan terlebih dahulu dalam function yang akan di defer

func endThisApp() {
	
	fmt.Println("End APP")
	
	message := recover() // disini recover akan menangkap pesan yang dikirimkan oleh panic, serta membatalkan proses panic yang menghentikan program. Jika recover() ditaruh dalam function yang sama baik sesudah ataupun sesudah baris kode panic(). maka, fungsi recover tidak akan jalan. Krn itu cara yang salah
	fmt.Println("Terjadi Error", message)

}

func includeThisPanicInRecover(error bool) {
	defer endThisApp()
	if error {
		panic("Error")
	}
}

func main() {
	includeThisPanicInRecover(true)
}