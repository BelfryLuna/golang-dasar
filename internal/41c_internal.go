package internal

import "fmt"

func init() {
	fmt.Println("This for internal")  // dengan menambahkan atau mengimport package internal, maka otomatis baris kode dari function init ini akan dieksekusi
}