package main

// VARIABLE
// variabel di golang adalah tempat menyimpan data

import "fmt"

func main() {
	// dalam penulisan variabel di golang. Kita menggunakan keyword var. Karena golang adalah static dan strong type language. saat deklarasi variabelnya kita harus menuliskan tipe data dari variabelnya.
	//  1, cara pertama, var nama_variabel tipe_data 
	var name string
	name = "Sonny Irsan" // disini adalah cara mengisi data dari variabel tersebuat (assign)
	fmt.Println(name)
	name = "eko" // disini kita hanya mengisi dan merubah data dalam variabelnya hanya dengan tipe data yang sesuai dari yang telah ditetapkan saat deklarasi
	fmt.Println(name)

	var name2 string = "Ika" // kita juga dapat mengisi data variabel secara langsung
	fmt.Println(name2)

	// 2, cara kedua. golang dapat secara langsung mengenali tipe data saat kita secara langsung menginisialisasi data dr variabelnya. sehingga kita tidak perlu lagi menuliskan tipe data saat deklarasi varibelnya dgn syarat nilainya langsung diisi
	
	var ass = 12 // seperti ini

	fmt.Println("#2", ass)

	// 3, cara ketiga. Go juga mengadopsi konsep type inference, yaitu metode deklarasi variabel yang tipe data-nya ditentukan oleh tipe data nilainya, cara kontradiktif jika dibandingkan dengan cara pertama. Dengan metode jenis ini, keyword var dan tipe data tidak perlu ditulis. jadi saat deklarasi variabelnya, nama_varibel := "valuenya". syarat: harus langsung diinisialisasi nilainya
	nama := "#3 Sonny"

	fmt.Println(nama)

	// 4, cara keempat, mendeklrasikan beberapa variabel sekaligus. syarat: harus langsung diinisialisasi nilainya
	var (
		namaDepan = "Sonny"
		namaBelakang = "Irsan" 
	)

	var namaTengah, namaSamping string

	namaTengah = "Irsan"
	namaSamping = "Tak Ada"

	var nama3, nama4 = "ini", "nama"

	nama1, nama2 := "Sonny", "Irsan"

	fmt.Println(nama3, nama4)
	fmt.Println(nama1, nama2)
	fmt.Println(namaTengah)
	fmt.Println(namaSamping)
	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)

	// 5, variabel underscore. krn golang mengharuskan tiap variabel yang dideklrasi harus digunakan. jadi ketika kita ingin membuat sebuah variabel yang tidak digunakan. dapat menggunakan variabel underscore (reserved variable)
	_ = "Sonny"
	// jadi meskipun tidak digunakan, data variabel ini tidak akan error

	// 6, variabel pointer, kita dapat membuat sebuah variabel menjadi pointer dengan keyword new
	iniPointer := new(string)
	iniPointer2 := iniPointer
	*iniPointer2 = "Sonny"

	fmt.Println("ini data pointer", *iniPointer2 , *iniPointer)

}