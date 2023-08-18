package main

import "fmt"

func main() {
	var buah = []string{"Apel", "jeruk", "mangga", "Durian", "Pisang"}
	var buah1 = buah[0:3]
	var buah2 = buah[0:2:2]

	fmt.Println("=======BUAH=========")
	fmt.Println(buah)
	fmt.Println(len(buah))
	fmt.Println(cap(buah))
	fmt.Println("=======BUAH1=========")
	fmt.Println(buah1)
	fmt.Println(len(buah1))
	fmt.Println(cap(buah1))
	fmt.Println("=======BUAH2=========")
	fmt.Println(buah2)
	fmt.Println(len(buah2))
	fmt.Println(cap(buah2))

}
