package main

import "fmt"

func main() {
	// pada Go tipe data number dibagi ke dalam dua jenis,
	// ...
	// integer(bilangan bulat) dan floating point(bilangan desimal)
	// pada integer(int), dibagi lagi ke tingkat yang lebih spesifik.
	// int8 -120an ke 120an, int16 -32ribuan ke 32ribuan, int32 -6milyaran ke 6milyaran, int64 lebih besar lagi
	// uint (jika tidak ingin negatif), sama uint8 s/d uint64
	// ...
	// floating point, juga dibagi lagi,
	// float32 dan float64 perbedaannya terdapat pada kemampuan menyimpan angka
    var num int 

	num = 13
	fmt.Println("ini number =", num)

	var floatIni float32 = 13.5
	fmt.Println("Ini float =", floatIni)
}