package main

import "fmt"


// Sama seperti di JS kita juga dapat membuat sebuah anonymous function di golang
// anonymous function adalah cara lebih singkat untuk menyimpan secara langsung sebuah function dalam variabel atau paramter (function tanpa nama)

type Blocked func(string) bool

func blockedName(name string, blacklist Blocked) {
	if blacklist (name) {
		fmt.Println("Anda diblacklist", name)
	} else {
		fmt.Println("Welcome," , name)
	}
}



func main() {
	// anonymous function dalam variabel (META)
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	blockedName("anjing", blacklist)

	// anonymous function dalam parameter, ditulis langsung functionnya tanpa sebuah nama
	blockedName("Eko", func(name string) bool {
		return name == "anjing"
		})

	// test
	tesAnonymous := func (a, b int) int  {
		return a + b
	}

	fmt.Println(tesAnonymous(12, 30))
}