package main

import "fmt"

func main() {


	// Hasil dari operasi perbandingan adalah berupa boolean, yaitu true dan false. 
	var banding = 32 != 10  // misalnya untuk operator != / tidak sama dengan, disini apakah 32 tidak sama dengan 10. jika benar atau iya kembalikan true. Dan jika tidak kembalikan false

	fmt.Println(banding) // krn perbandingan benar, disini akan mengemablikan nilai true``

	var tesStr = "a" == "b"

	fmt.Println(tesStr)

	var angka1 = "perempuanaaa"
	var text1 = "perempuana"

	fmt.Println("hasilnya , " , angka1 != text1)


	perbandinganNilai := 12 > 7

	fmt.Println(perbandinganNilai)

}