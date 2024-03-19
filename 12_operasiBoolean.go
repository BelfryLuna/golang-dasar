package main

import "fmt"

func main() {
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