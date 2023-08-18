package main

import "fmt"

func main() {
	var buah = []string{"Apel", "jeruk", "mangga", "Durian"}
	fmt.Println(buah[0:2])
	fmt.Println(buah[0:4])
	fmt.Println(buah[0:0])
	fmt.Println(buah[4:4])
	//fmt.Println(buah[4:0]) error
	fmt.Println(buah[:])
	fmt.Println(buah[2:])
	fmt.Println(buah[:2])
}
