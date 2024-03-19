package main

import "fmt"

func main() {

	// 1
	var myArray = [...]string{
		"M. Sonny",
		"Irsan",
		"Syahputra",
	}

	fmt.Println(myArray[0])


	// 2 
	var myArray1 [5]string

	myArray1[0] = "M. Sonny"
	myArray1[1] = "Irsan"

	fmt.Println(myArray1)

	// 3
	var myArray2 = [4]int {1,2,3}
	
	fmt.Println(myArray2)

	// 4
	var myArray3 = [2]string {
		"Sonny",
		"Irsan",
	}

	fmt.Println(myArray3)

	var testArray = [6]string {
		"Sonny",
		"Irsan",
		"syahputra",
		"Sonny",
		"Sonny",
		"Sonny",
	}

	fmt.Println(len(testArray))

	// array multidimensi
	arrayMulti := [3] [3]int{{1,2,3} , {4,5,6}, {7,8,9}}
	fmt.Println("array indeks ke 1", arrayMulti[1])

	// membuat array dengan make
	arrayMake := make([]string, 2)

}