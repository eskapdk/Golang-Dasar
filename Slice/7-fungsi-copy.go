package main

import "fmt"

func main() {
	var tujuan = make([]string, 3)
	var sumber = []string{"Jeruk", "Apel", "jambu", "Pisang"}
	var cpy = copy(tujuan, sumber)

	fmt.Println(tujuan)
	fmt.Println(sumber)
	fmt.Println(cpy)

	fmt.Println("==================================")
	var tujuan2 = []string{"Apel", "Apel", "Apel"}
	var sumber2 = []string{"Jeruk", "Jeruk", "Jeruk", "Jeruk"}
	var cpy2 = copy(tujuan2, sumber2)

	fmt.Println(tujuan2)
	fmt.Println(sumber2)
	fmt.Println(cpy2)

}
