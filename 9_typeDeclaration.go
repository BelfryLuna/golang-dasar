package main

import "fmt"

func main() {
	// ingin membuat sebuah tipe data yang menampung nilai yang aslinya number, tapi kita dapat membuat sebuah
	// tipe data baru/alias jika ingin menyimoan data no ktp sbg string
	// jadi kita dapat membuat tipe data alias dari tipe data string untuk merepresentasikan nilai dari no ktp yang mulanya number ke string

	type noKtp string  // mekanismenya seperti kita memberikan alias untuk tipe data tertentu dgn tujuan mempermudah, disini kita membuat tipe data baru yaitu noKTP, yang sebenarnya merupakan string. jadi kita membuat alias baru untu tipe data string dgn nama "noKtp"

	var noKtpSonny noKtp = "1407020405010005"

	fmt.Println(noKtpSonny)

	var noKtpAsli string = "1407020405010005"

	var konversi noKtp = noKtp(noKtpAsli)

	fmt.Println(noKtpAsli)
}