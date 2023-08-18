package main

import "fmt"

func main() {
	var buah = [4]string{"apel", "jeruk", "nanas", "melon"}
	for _, element := range buah {

		fmt.Printf("nama buah : %s\n", element)
	}
}
