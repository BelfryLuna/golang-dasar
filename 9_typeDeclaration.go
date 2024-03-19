package main

import "fmt"

func main() {
	// ingin membuat sebuah tipe data yang menampung nilai yang aslinya number, tapi kita dapat membuat sebuah
	// tipe data baru/alias jika ingin menyimoan data no ktp sbg string
	// jadi kita dapat membuat tipe data alias dari tipe data string untuk merepresentasikan nilai dari no ktp yang mulanya number ke string

	type noKtp string

	var noKtpSonny noKtp = "1407020405010005"

	fmt.Println(noKtpSonny)

	var noKtpAsli string = "1407020405010005"

	var konversi noKtp = noKtp(noKtpAsli)

	fmt.Println(noKtpAsli)
}