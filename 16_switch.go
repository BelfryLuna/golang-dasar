package main

import "fmt"

func main() {
	var nama string = "Sonny"

	// switch adalah versi sederhana dari if, dimana kita tetap melakukan pengondisian. tetapi, hanya terhadap kondis pada satu variabel saja
	// misalnya
	switch nama { // kita melakukan pengondisian yang beragntung dari isi dari satu variabel nama
	case "Sonny": // disini if diaganti dengan case (kasus), jika case atau isi dari variabel nya adalah "sony" dan memenuhi kondisi case tersebut
		fmt.Println("Hello, ", nama) // maka, jalankan ini
	case "Eko": // disini adalah case2 berikutnya yang tak ubahnya else if
		fmt.Println("Halo, ", nama)
	default: // default berperan sbg else disini
		fmt.Println("Hello, boleh kenalan?")
	}

	// jika pengondisian hanya melibatkan kasus yang sederhana atau melibatkan satu bariabel maka gunakan switch

	// switch dgn short statement
	switch name := "Eko"; name {
	case "Sonny": // disini if diaganti dengan case (kasus), jika case atau isi dari variabel nya adalah "sony" dan memenuhi kondisi case tersebut
		fmt.Println("Hello, ", name) // maka, jalankan ini
	case "Eko": // disini adalah case2 berikutnya yang tak ubahnya else if
		fmt.Println("Halo, ", name)
	default: // default berperan sbg else disini
		fmt.Println("Hello, boleh kenalan?")
	}

	// switch tanpa kondisi, jadi kondisinya ditaruh pada tiap2 casenya (mirip if else)

	umur := 35

	switch {
	case umur > 40:
		fmt.Println("Sudah tua")
	case umur > 17:
		fmt.Println("Sudah dewasa")
	default:
		fmt.Println("Belum Dewasa")
	}

	// switch dengan banyak case

	nilai := "A"

	switch nilai {
	case "A", "B":
		fmt.Println("Sangat Memuaskan")
		fallthrough // ini digunakan untuk memaksa switch, untuk terus mengeksekusi case setelahnya (meskipun true atau false casenya).
	case "C", "D":
		fmt.Println("Lulus")
	default:
		fmt.Println("Tidak Lulus")
	}
}
