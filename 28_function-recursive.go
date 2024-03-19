package main

import "fmt"

func faktorialLoop(value int) {
	temp := 1

	for i := value; i > 0; i-- {
		temp *= i
	}

	fmt.Println(temp)
}

func recursiveFactorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursiveFactorial(value - 1)
		// 5 * recursiveFactorial(5 - 1)
		// 5 * 4 * recursiveFactorial(4 - 1)
		// 5 * 4 * 3 * recursiveFactorial(3 - 1)
		// 5 * 4 * 3 * 2 * recursiveFactorial(2 - 1)
		// 5 * 4 * 3 * 2 * 1 recursiveFactorial(1 - 1)
		// 5 * 4 * 3 * 2 * 1 (ketika semua function telah terpanggil barulah dijalankan faktorialnya)
	}
}


func bilPrima (num int) bool {
	if num == 1 {
		return false
	}

	for i := 2; i < num; i++ {
		if num % i == 0 {
			return false
		}
	}

	return true
}



func main() {
	newSlice := []int {3, 5, 6, 7, 11, 17, 19, 20}

	arrayNew := [7]int {}
	for index, elemen := range newSlice {
		if bilPrima(elemen)  {
			arrayNew[index] = elemen
		} 
	}
	fmt.Println(arrayNew)

	
	


	
}