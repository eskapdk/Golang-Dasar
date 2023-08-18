package main

import "fmt"

func main() {
	var buah = []string{"Apel", "jeruk", "mangga", "Durian", "Pisang"}
	var buah1 = buah[0:3]
	var tambah = append(buah1, "Nanas")
	var tambah2 = append(tambah, "Duku", "Maggis")

	fmt.Println(buah)
	fmt.Println(buah1)
	fmt.Println(tambah)
	fmt.Println(tambah2)
	fmt.Println("==========BUAH========")
	fmt.Println(len(buah))
	fmt.Println(cap(buah))
	fmt.Println("==========BUAH1========")
	fmt.Println(len(buah1))
	fmt.Println(cap(buah1))
	fmt.Println("==========TAMBAH========")
	fmt.Println(len(tambah))
	fmt.Println(cap(tambah))
	fmt.Println("==========TAMBAH2========")
	fmt.Println(len(tambah2))
	fmt.Println(cap(tambah2))

}
