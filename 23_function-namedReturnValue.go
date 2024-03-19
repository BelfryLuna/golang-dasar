package main

import "fmt"

// Kita juga dapat memberikan nama(ibarat sebuah variabel) terhadap return value nya. dengan memberikan nama pada saat membuat tipe data pengembalian di sebuah function

// dan named return value tersebut juga dapat diisi layaknya variabel


func namedRetur() (fisrtName string, lastName string) {
	fisrtName = "Sonny" //named return value dapat diisi layaknya variabel
	lastName = "Irsan"

	return fisrtName, lastName // dan disini kita dapat menuliskan named return nya saja pada return
}

func coba() (firt int, last string) {
	firt = 22
	last = "Sonny"

	return firt, last // return value diurutkan sesuai dengan dengan urutan saat membuat tipe data pengembalian
}

func main() {

	a, b := coba()

	fmt.Println(a,b)

	tryFirstName, tryLastName := namedRetur()

	tryAgain, _ := namedRetur()

	fmt.Println(tryAgain)

	fmt.Println(tryFirstName, tryLastName)
}