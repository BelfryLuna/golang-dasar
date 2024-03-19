package main

import "fmt"

func main() {
	var ini32 int32 = 32768

	var ke64 = int64(ini32)

	var result = ke64

	fmt.Println(result)


	str := "Sonny"
	ambilStr := str[1]
	keStr := string(ambilStr)

	fmt.Println(keStr)


	var tes16 int16 = 32768

	fmt.Println(tes16)

}