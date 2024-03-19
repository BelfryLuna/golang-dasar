package main

import "fmt"

// closure merupakan sebuah fitur dimana kita dapat mengakses variabel serta merubah didalam sebuah scope(function) yang berbeda 
// atau, kemampuan sebuah function untuk berinteraksi dgn data2 di sekitarnya (data/variabel/function dalam global scope atau di dalam main function)

func cobalah() {
	iniVar := 0
	fmt.Println(iniVar)
}


func main() {
	count := 0
	count2 := 5


	increment := func ()  {
		fmt.Println("Increment+")
		count++
	}

	decrement := func ()  {
		fmt.Println("Decrement-")
		count2--
	}


	increment()
	increment()

	decrement()
	decrement()

	fmt.Println(count2)
	fmt.Println(count)
}