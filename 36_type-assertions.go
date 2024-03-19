package main

import "fmt"

// TYPE ASSERTIONta 
// kemampuan di golang yang digunakan untuk merubah tipe data dari sebuah value yang seringnya diinisialisasi pada interface kosong menjadi tipe data yang kita inginkan.

// misalnya, kita membuat function yang mengembalikan any/interface
func randomValue() any {
	return "oke"
}

func main() {
	// ketika kita simpan datanya pada variabel, maka variabel tersebut dianggap tipe datanya adalah any. bukan string, meskipun return dr function tersebut adalah string.
	result := randomValue()

	// kita dapat mengubah tipe data yang any tersebut menjadi string, menggunakan type assertion.
	var resultStr string
		resultStr = result.(string) // result.(tipeData_yangDiinginkan) merupakan syntax jika ingin melakukan type assertion
		fmt.Println(resultStr) // maka disini variabel resultStr tidak lagi any, melainkan string

	// kenapa tidak dapat menggunakan konversi tipe data biasa, krn akan error sehingga dibutuhkan type assertions
	// var konversiBiasa string = string(result) 

	// namun, menjadi perhatian jika nilai pada return valuenya, berbeda dr tipe data yang diinginkan. cth pada function yng mengembalikan any, pada return nya kita mengisi nilai string. ketika kita ingin menggunakan type assertions menjadi integer, maka hal ini akan menghasilkan panic. karena, return nya adallah string, jadi tidak bisa diubah menjadi integer meskipun menggunakan type assertions. jadi harus hati-hati.
	var resultInt int
		resultInt = result.(int)
		fmt.Println(resultInt)

	// untuk tidak terjadi panic, kita dapat menggunakan fitur di swutch yaitu type,
	// type, dapet menangkap apa tipe data return value dr function any tsb
	// sehingga dapat dilakukan percabangan, jadi ketika tipe datanya adalah string misalnya, maka akan melakukan type assertions yang sesuai. dan jika integer, maka lakukan type assertion yang sesuai dgn integer pula.
	switch value := result.(type) { // akan disimpan dalam var value dulu, barulah diisi dgn result.(type) (syntax fitur typenya switch)
	case string: // jadi switch ini dapat mengenali, dan mencabangkan eksekusi per case (tipe datanya). jika type nya string lakukan di bawah
		resultStr = result.(string)
			fmt.Println("string", resultStr)
	case int: // jika integer, eksekusi baris kode di bawah
		resultInt = result.(int)
			fmt.Println("int", resultInt)
	default: 
		fmt.Println("Unknown", value)
	}
	// hal ini akan membuat program hanya akan menjalankan type assertion yang sesuai dan menghindari panic
}