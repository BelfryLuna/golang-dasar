package main

import "fmt"

func bilGenap(num int) {
	for i := 0; i <= num; i++ {
		if i %2 == 0 {
			fmt.Println("Perulangan ke", i)
		}
	}
}


func main() {
	for i := 0; i < 15; i++ {
		if i%2 == 1 {
			continue  // fungsi continue pada perulangan berguna untuk menghentikan perulangan yg skrg (ngeskip) dan melanjutkannya pada perulangan berikutnya (jadi i yang bernilai ganjil akan di skip)
		}

		if i == 14 {
			break // akan menghentikan keseluruhan proses perulangan
		}

		fmt.Println("Perulangan ke, " , i)
	}

	bilGenap(20)

}