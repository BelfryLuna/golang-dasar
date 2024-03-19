package main

import "fmt"

// NIL
// #1 pada golang, bila kita membuat variabel dan serta telah mendeclare tipe datanya. ketika kita tidak mengisi datanya dan mencoba mengaksesnya dengan Println() secara otomatis akan terisi oleh default value dr tipe datanya

// memberikan nilai nil
// contoh pada map, 
func newMap(nama string) map[string]string {
	if nama == "" {
		return nil
	} else {
		return map[string]string{
			"name" : nama,
		}
	}
} 

// pada any
func dataInterfaceKosongTapiNil(nama any) any {
	if nama == "" {
		return nil
	} else {
		return nama
	}
}

// pada slice
func makeSice(nomor int) []int {
	if nomor == 0 {
		return nil
	} else {
		return []int{
			nomor,
		}
	}
}

func main() {
	// #2 disini kita mencoba mendeclare variabel dan tipe datanya, lalu Println() semua variabelnya tanpa mengisi valuenya
	var str string
		fmt.Println(str) // ""

	var intg int
		fmt.Println(intg) // 0
	
	var boole bool
		fmt.Println(boole) // false
	
	// #3 ketiga variabel tsb mnegembalikan default value masing2 tipe datanya
		// #4 berbeda dgn javascript, hal ini akan berujung undefined atau null di bhs lain

// di Golang juga terdapat hal tersebut, tapi dinamakan nil. nil dapat diimplementasikan pada map, function, slice, ponter dan channel, interface. selain ini maka nilai nil tidak bisa diimplementasikan.

	// disini kita dapat melihat apakah sebuah map tersebut benar2 memilki nilai nil dgn if else
	fmt.Println(newMap(""))
		dataNil := newMap("")
		if dataNil == nil {
			fmt.Println("ya ini nil")
		} else {
			fmt.Println("Bukan ini mah")
		}

		dataInterfaceNill := dataInterfaceKosongTapiNil("")
		if dataInterfaceNill == nil {
			fmt.Println("ya ini nil")
		} else {
			fmt.Println("Bukan ini mah")
		}
	
	fmt.Println(makeSice(0))
		dataSlice := makeSice(0)
			if dataSlice == nil {
				fmt.Println("kita bersama nil")
			} else {
				fmt.Println(dataSlice)
			}

	// coba mengisi nilai nil secara langsung di variabel

	// pada interface kosong
	var dataNilIni any = nil

	// pada map
	var mapping map[int]string
	mapping = nil

	// pada slice
	newSlice := make([]string, 2, 5)
	newSlice = nil

	if dataNilIni == nil && mapping == nil && newSlice == nil{
		fmt.Println("ya data masih kosong")
	} else {
		fmt.Println("0")
	}


	// Jadi, ketika ingin membuat sebuah data kosong, maka gunakan nil

}