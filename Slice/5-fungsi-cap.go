package main

import "fmt"

func main() {
	var buah = []string{"Apel", "jeruk", "mangga", "Durian", "Pisang"}
	fmt.Println(len(buah))
	fmt.Println(cap(buah))

	fmt.Println("=========SLICE 1========")
	var buah1 = buah[0:3]
	fmt.Println(len(buah1))
	fmt.Println(cap(buah1))

	fmt.Println("=========SLICE 2========")
	var buah2 = buah[1:4]
	fmt.Println(len(buah2))
	fmt.Println(cap(buah2))
}
