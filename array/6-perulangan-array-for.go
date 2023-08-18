package main

import "fmt"

func main() {
	var buah = [4]string{"apel", "jeruk", "nanas", "melon"}
	for i := 0; i < len(buah); i++ {
		fmt.Println("Isi Array Buah", i, ":", buah[i])
	}
}
