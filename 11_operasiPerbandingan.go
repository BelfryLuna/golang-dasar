package main

import "fmt"

func main() {


	// Hasil dari operasi perbandingan adalah berupa boolean, yaitu true dan false
	var banding = 32 != 10

	fmt.Println(banding)

	var tesStr = "a" == "b"

	fmt.Println(tesStr)

	var angka1 = "perempuanaaa"
	var text1 = "perempuana"

	fmt.Println("hasilnya , " , angka1 != text1)

}