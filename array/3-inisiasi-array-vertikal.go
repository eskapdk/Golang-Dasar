package main

import "fmt"

func main() {
	//cara horisontal
	var buah = [4]string{"apel", "jeruk", "nanas", "melon"}
	//cara vertikal
	var buah2 = [4]string{
		"Semangka",
		"Anggur",
		"Pisang",
		"Pir",
	}

	fmt.Println("Panjang Array \t\t", len(buah))
	fmt.Println("Isi Array \t\t", buah)

	fmt.Println("Panjang Array \t\t", len(buah2))
	fmt.Println("Isi Array \t\t", buah2)
}
