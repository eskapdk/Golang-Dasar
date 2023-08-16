package main

import "fmt"

func main() {
	var nilai = 7
	switch {
	case nilai == 10:
		fmt.Println("Sempurna")
	case (nilai < 10) && (nilai > 6):
		fmt.Println("Bagus")
	case (nilai < 7) && (nilai > 4):
		fmt.Println("Lumayan")
	default:
		fmt.Println("Kurang")
		fmt.Println("Coba lagi")
	}
}
