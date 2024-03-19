package main

import "fmt"

// Custom Error

/*
Kita dapat membuat custom error kita sendiri dengan cara mengikuti kontrak dari interface bawaan error di go lang
caranya:
kita buat terlebih dahulu sebuah struct yang memiliki sebuah method yang mengikuti kontrak method dari interface nya error bawaan go lang (membuat struct yang mengikuti kontrak dari interface error bawaan go lang)
*/

// pertama buat dulu struct nya
type ValidationError struct {
	Massage string
}

// kemudian bua sebuah method yang mengikuti kontrak dr interface bawaan golang
func (v *ValidationError) Error() string  {
	return v.Massage
}

// nah di atas kita telah berhasil membuat sebuah kustom error dengan mengikuti struktur/kontrak dari interface error bawaan go lang

type NotFoundUserError struct {
	Massage string
}

// seperti yang sebelumnya, wajib sekali menjadikan ethod dalam pointer
func (n *NotFoundUserError) Error() string  {
	return n.Massage
}

func saveData(id string, data any) error {
	if id == "" {
		return &ValidationError{Massage: "Validation Error"} // krn tipe data kembaliannya adalah error yang merupakan interface jadi kita perlu mengembalikannya dalam pointer
	} else if id != "eko" {
		return &NotFoundUserError{Massage: "Not Found User Error"}
	} 

	return nil
}

func main() {
	// disini kita akan mencari tahu detail dr sebuah error apakah error tsb termasuk salah satu jenis2 custom error yang telah dibuat sebelumnya atau bukan
	getErr := saveData("", nil)

	if getErr != nil {
		// Jika Error
		
		/* variabel cekError akan menangkap tipe error dari getErr (dgn type assertion) */
		switch cekError := getErr.(type)  {
		case *ValidationError: // jika tipe dari error nya adalah *ValidationError, lakukan di bawah
			fmt.Println("Error :", cekError.Error())
		case *NotFoundUserError:
			fmt.Println("Error :", cekError.Error())
		default:
			fmt.Println("Error :",  cekError.Error())
		}
	} else {
		// Jika Sukses
		fmt.Println("Sukses")
	}
}