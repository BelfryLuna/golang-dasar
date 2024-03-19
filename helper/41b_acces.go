package helper

// ACCES MODIFIER

// Kita dapat memodifikasi acces dari setiap variabel, function dll dalam golang, hal ini berlaku ketka:
// data variabel atau function tsb hendak digunakan atau diakses dr luar package nya.

/* di GOLANG
		dalam membuat acces modifier terhadap sebuah data kita hanya perlu mengubah huruf awal dr nama masing2 data tersebut.
		misalnya,
		jika ingin data tersebut diakses di luar packagenya:
		kita hanya perlu membuat nama dari data (function, variable dll) tersebut dr huruf kapitas
		cth: - Var Nama string
			 - func SayHello() {
					fmt.Println("Hello World")
			   }

		Jika tidak ingin data tersebut dapat diakses oleh package lain (namun, masih dapat diakses di dalam packagenya sediri)
		kita hanya perlu membuat data (variabel, function dll) seperti biasa.
		cth: - var nama string
			 - func sayHello() {
					fmt.Println("Hello World")
			   }
*/

// contoh dapat diakses oleh package lain 

var Namae string = "Sonny"

func SayHi (nama string) string {
	return "Hai " + nama 
}

// tidak dapat diakses (tapi dapat diakses dalam package yang sama)

var named string = "Sonny"

func sayHello(nama string) string  {
	return "Hai " + nama
}