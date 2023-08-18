package main

import "fmt"

func main() {
	var buah = []string{"Apel", "jeruk", "mangga", "Durian", "Pisang"}

	var buah1 = buah[0:3]
	var buah2 = buah[1:4]

	var buah11 = buah1[0:1]
	var buah22 = buah2[0:2]

	fmt.Println(buah)
	fmt.Println(buah1)
	fmt.Println(buah2)
	fmt.Println(buah11)
	fmt.Println(buah22)

	fmt.Println("========DIRUBAH=======")
	buah22[0] = "Jambu"

	fmt.Println(buah)
	fmt.Println(buah1)
	fmt.Println(buah2)
	fmt.Println(buah11)
	fmt.Println(buah22)
}
