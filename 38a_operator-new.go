package main

import "fmt"

// OPERATOR NEW (META, untuk mempermudah banyak programmer menggunakan new dr pada operator & jika ingin men=mbuat sebuah pointer kosong)
// fungsi new() dapat digunakan untuk membuat sebuah pointer (akan dibuat sebuah pointer menuju ke memory value baru yang datanya masih kosong)

// new(memoryValueBaru)

type Alamat struct {
	Kota, Provinsi, Negara string
}

func main() {

	alamat1 := new(Alamat) // disini akan membuat pointer yang menuju value baru untuk template/struct Alamat, maka ketika menggunakan opearator new dalam sebuah variabel, maka variabel tersebut akan menjadi senuah pointer
	alamat2 := alamat1

	alamat2.Kota = "Pekanbaru"
	
	fmt.Println(alamat1)
	fmt.Println(alamat2)

	*alamat2 = Alamat {"Pekanbaru", "Riau", "Indonesia"}

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}