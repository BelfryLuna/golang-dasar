package main

import "fmt"

// OPERASI BOOLEAN TERDIRI DARI
/*
	&&     = Dan,
	akan mengemablikan true jika kedua/kesemua operan/data yang dilogikakan(boolean kan) bernilai true
	misal,    12 > 7 && 12 == 12
							true
		      
			  12 < 3 && 12 != 12
			  				false

	||     = atau
	akan mengembalikan nilai true, jika salah satu saja dari operannya bernilai true
	misal,    12 < 8 || 12 == 12
							true

	!data   = kebalikan
	jika data yang awalnya adalah false, jika menambahkan ! di awal akan emnjadikan data menjadi kebalikannya
	misalnya
*/

func main() {

	and := 12 < 7 && 12 == 12
	fmt.Println(" and", and)  // kesemua operan tidak bernilai true krn operan pertama adalah false dan operan kedua true, maka hasilnya false

	or := 12 < 7 || 12 == 12
	fmt.Println("2 or", or) // krn salah satu nilai operannya true, maka hasilnya akan mengembalikan nilai true

	fmt.Println("kebalikan", !or) // disini variabel or yang awalnya true, nilainya akan berubah menhjadi flase dgn operator ! di awal

	var (
		nilaiUjian   = 75
		absensi = 80
	)

	

	var (
		akhirUjian   = nilaiUjian >= 70
		akhirAbsensi = absensi > 80
	)

	fmt.Println("Lulus ujian?", akhirUjian)
	fmt.Println("lulus absensi?", akhirAbsensi)

	var hasil = !akhirAbsensi

	fmt.Println("Apakah lulus? ", hasil)
}