package main

import "fmt"

// INTERFACE (tipe data)
// dalam golang, interface adalah sebuah tipe data abstrak.

// interface ini berisi kumpulan definisi2 method, sbg kontrak atau struktur yang nantinya akan diimplementasi kan berdasarkan kontrak dr method2 yang telah didefinisikan di interface

// dalam pembuatan sebuah interface, caranya hampir sama seperti struct
// menggunakan type declaration diikuti nama_interface dan keyword interface, lalu diisi dgn kontrak dr method-mthod (struktur dr methodnya) yang ingin diisi dalam interface

type HasName interface { // type nama_interface interface
	getName() string // didalamnya berisi kontrak method (struktur2nya, dr nama method, ada atau tidaknya parameter dan tipe data pengembalian jika ingin menambahkan return value)

	// nah nanti ketika kita ingin mengimplementasikan method dalam interface tsb, kita harus mengisinya dgn method dgn struktur yang serupa dari nama dst...
}



// pengimplementasian interface
// disini krn interface juga dikenali sbg tipe data, kita dapat menjadikannya sbg tipe data dri variabel atau parameter misalnya,
func sayHai(hasName HasName) /*nah data yang harus dikirimkan dalam parameter ini harus merupakan sebuah method (dari struct) yang memiliki struktur yang sama dengan kontrak yang telah ditetapkan dalam interface tersebut*/ {
	fmt.Println("Hai", hasName.getName() + ", Bagus nih")
}


// sebagaimana yang telah dijelaskan, kita dapat mengimplementasikan data dari interface sebelumnya kita harus membuat struct dahulu dan membuat method,
type GotName struct {
	Nama string
}

// disini kita akan membuat sebuah mthod dr GotName struct yang nantinya akan diimplementasikan sbg data yang akan dikirimkan pada interface,
// dimana, kita harus membuat method tsb memiliki struktur yang sama (dari nama method dst...) dgn kontrak method yang didefinisikan pada interface
func (gotName GotName) getName() string {
	return gotName.Nama
}


// mencoba mengisi oarameter interface dgn function dgn struktur yg sama dengan kontrak (ERROR), krn interface hanya menerima sebuah method
func getName() string {
	return "Halo"
}

// kita juga dapat membuat bnyk struct (asal memiliki method yg sesuai dgn kontrak di interface) untuk mengimplementasikan data dr satu interface

type Car struct {
	Name string
	Rilis int
}

func (car Car) getName() string {
	return car.Name
}

func main() {
	// dalam memanggil sebuah function (sayHai) yang menjadikan interface sbg parameternya, data yang kita kirimkan berupa,

	// disini kita akan membuat object yang mengisi data pada struct
	person := GotName {
		Nama: "Sonny",
	}
	// krn variabel/objek person telah memiliki method dgn struktur yang sama dgn kontrak, maka var person dapat dikirimkan sbg data dri parameter interface teresebut
	sayHai(person) // disini person menyimpan field dan juga method (sesuai kontrak), makanya tidak error


	animal := Car {
		Name: "Avanza",
		Rilis: 2014,
	}

	sayHai(animal)

	fmt.Println(person.getName())












	// NOTE: perlu diingat bahwa method dan function berbeda di golang, yang bisa mengimplementasikan data interface haruslah sebuah method, meskipun kita memiliki sebuah function yg sama dgn kontrak, tapi akan menimbulkan error

	// sayHai(getNameThis())
}