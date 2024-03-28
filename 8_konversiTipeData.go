package main

import "fmt"

func main() {
	// konversi tipe data biasa (binary ke string, sesama number atau float misal int8 ke int32 dst)dgn penulisan lengkap
	var ini32 int32 = 32768

	var ke64 = int64(ini32)

	var result = ke64

	fmt.Println(result)


	// cara singkat
	str := "Sonny"
	ambilStr := str[1]
	keStr := string(ambilStr) // caranya ketikkan tioe data yang diinginkan string(data_yangInginDiubah)

	fmt.Println(keStr)


	var ini8 int8 = 18
	ke32 := int32(ini8)

		fmt.Println()

}