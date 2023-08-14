package main

import "fmt"

func main() {

	//deklarasi variabel sama
	var satu, dua, tiga string
	satu, dua, tiga = "one", "two", "three"

	var empat, lima, enam string = "four", "five", "six"
	tujuh, delapan, sembilan := "seven", "eight", "nine"

	//deklarasi beda variable
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	fmt.Println(satu, dua, tiga, empat, lima, enam, tujuh, delapan, sembilan)
	fmt.Println(one, isFriday, twoPointTwo, say)
}
