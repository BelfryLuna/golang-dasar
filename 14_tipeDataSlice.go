package main

import "fmt"

func main() {
	// tipe data slice (META)
	// potongan dari array, slice dapat dibuat melalui mengakses sebagian atau seluruh data dari array
	// slice ukurannya dapat berubah

	// komponen dalam slice
	// pointer, penunjuk data pertama dari data array yang dijadikan untuk data pertama di slice
	// length, panjang dari slice
	// capacity, kapasitas dari slice nya

	// membuat slice dari array

	// sampel array yang akan dijadikan slice
	var ambilArray = [7]int{
		1, // index 0
		2, // index 1
		3, // index 2  , pointer pada mySlice
		4, // index 3
		5, // index 4
		6, // index 5 , batas data dr mySlice, lenght : dari data index 2 ke index 5

		7, // index 6 , batas capacity dr mySlice, dari data pointer ke akhir array
	}

	mySlice := ambilArray[2:6]

	mySlice1 := ambilArray[3:] // mengambil data array, dari pointer di index 3 sampe index terakhir di array

	mySlice2 := ambilArray[:4] // menganmbil data array dari index pertama array sampai sebelum index height (index 3)

	mySlice3 := ambilArray[:] // mengambil seluruh data pada array

	fmt.Println(mySlice3)

	fmt.Println(mySlice1) // 4 5 6 7

	fmt.Println(mySlice2) // 1 2 3 4

	fmt.Println(mySlice) // output : [3 4 5 6]

	//  mySlice := ambilArray[:]  sama saja var mySlice []string = ambilArray[:]
	// jika pada array kita harus mengisi jumlah data yang diisi atau memberi titik tiga jika ragu,
	// [3]string / [...]int pada slice tidak perlu cukup []string / []int

	// build function khusus slice

	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice2))

	// fungsi append di slice, menambahkan data baru di array, data akan ditambahkan di akhir serta akan ditambahkan di array 
	// jika belum melebihi kapasitas, jika memlebihi kapasitas maka data yang baru tersebut otomatis dibuatkan
	// array baru oleh slice untuk menampung data tersebut


	days := [...]string {"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}

	daySlice1 := days[3:7]

	daySlice1[0] = "Thursday" // disini saat kita mengubah isi data dari slice sama halnya kita juga mengubah data dari array nya
	daySlice1[1] = "Friday"
	fmt.Println(daySlice1)
	fmt.Println(days)

	// Append, ketika fungsi append digunakan maka akan menambahkan data baru di akhir. jika capacity slice dr array tidak lagi 
	// dapat menampung. maka akan dibuatkan sebuah array baru dalam internal logic si golangnya 
	// dan juga pada saat menggunakan fungsi append kita otomatis akan membuat sebuah slice baru

	fmt.Println("Batas")

	daySlice2 := append(daySlice1, "Ahad") // disini data baru akan ditambahkan dari data dayslice1 dan dimasukkan ke dalam dayslice2, namun data yang diappend tadi tidak masuk ke dayslicee1
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	daySlice2[0] = "Kamis" // disini yang akan berubah hanyalah data dari dayslice2 dan array baru yang dibuat di dalam internal logic di dayslicee2
	fmt.Println(daySlice2)
	fmt.Println(daySlice1)
	fmt.Println(days)
	
	// membuat slice baru serta membuat array baru menggunakan fungsi make(keywordSlice, length dari slice, dan capacity(panjang) array)
	fmt.Println("Batas")

	newSlice := make([]string, 3, 5) // atau var newSlice []string = make([]string, 3, 5)
	// penjabaran
	// newSlice := make([]string, 3, 5)
	// 3, disini adalah panjang awal dari slice yang ingin dibuat (dapat berubah jika ingin mengappend (slice baru))
	/* 5, kapasitas dari array baru yang dibuat dengan fungsi make(), jadi selama kapasitasnya belum penuh, 
	   saat meng-append data baru akan disimpan di array tersebut. selagi kapasitasnya masih ada maka tidak akan dibuat
	   array baru. jika kapasitasnya sudah penuh, barulah dibuat array baru jika ingin meng-append data yang baru*/ 

	newSlice[0] = "M. Sonny"
	newSlice[1] = "Irsan"
	newSlice[2] = "Syahputra"

	fmt.Println(newSlice)

	newSlice2 := append(newSlice, "Bin") // disini data yang di append masih masuk ke dalam array pertama. namun data barunya hanya akan masuk ke newSlice2 dan length dari newSlice2 menjadi 4
	

	fmt.Println(newSlice2)

	fmt.Println("Batas append yang baru")
	newSlice2[0] = "Muhammad" // krn masih dari array yang sama maka, data index 0 dari newslice dan newslice2 akan smaa2 berubah krn masih dalam array yang sama
	fmt.Println("Total lenght newSlice2 = ", len(newSlice2))
	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	// men-copy slice, mengambil data slice a dan dicopy keseluruhan isi data dari slice a ke slice b

	day := [4]string {"Sonny", "Zaki", "Mama", "Ayah"}

	fromSlice := day[:]

	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	// disini fungsi copy() yang digunakan, copy(keSliceMana, dariSliceMana)

	fmt.Println("Batas lagi")
	fmt.Println(fromSlice)
	fmt.Println(toSlice)


	// Membuat slice dengan mengisi langsung datanya
	sliceNew := []string {"M. Sonny", "Irsan", "Syahputra"} // disini akan membuat sebuah slice baru dan membuat sebuah array baru

	fmt.Println(len(sliceNew), cap(sliceNew), sliceNew)

	sliceNew2 := append(sliceNew, "bin")

	sliceNew2[0] = "Muhammad"
	fmt.Println(sliceNew)
	fmt.Println(sliceNew2)

	// di golang yang lebih sering digunakan adalah slice daripada array

}