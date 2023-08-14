package main

import "fmt"

func main() {
	const firstName string = "eska"
	fmt.Print("Halo: ", firstName, "!\n")

	const lastName = "pratama"
	fmt.Print("Halo: ", lastName, "!\n")

	fmt.Println("===========")
	//perbedaan println dan print
	fmt.Println("Eska Pratama")
	fmt.Println("Eska", "Pratama")

	fmt.Print("Eska Pratama\n")
	fmt.Print("Eska", "Pratama\n")
	fmt.Print("Eska", " ", "Pratama\n")

	//konstanta multi
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	fmt.Println("===========")
	fmt.Println("Square :", square)
	fmt.Println("is Today :", isToday)
	fmt.Println("float :", floatNum)

	//konstanta tipe data sama
	const (
		a = "Konstanta"
		b
		today string = "senin"
		sekarang
		isToday2 = true
	)
	fmt.Println("===========")
	fmt.Println("a :", a)
	fmt.Println("b :", b)
	fmt.Println("today :", today)
	fmt.Println("sekarang :", sekarang)
	fmt.Println("is today 2 :", isToday2)

	//konstanta 1 baris
	const satu, dua = 1, 2
	const tiga, empat string = "tiga", "empat"

	fmt.Println(satu, dua, tiga, empat)

}
