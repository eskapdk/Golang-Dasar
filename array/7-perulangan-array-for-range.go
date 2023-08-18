package main

import "fmt"

func main() {
	var buah = [4]string{"apel", "jeruk", "nanas", "melon"}
	for i, elemet := range buah {
		fmt.Printf("Isi Array %d : %s\n", i, elemet)
	}
}
