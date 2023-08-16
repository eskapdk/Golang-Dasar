package main

import "fmt"

func main() {
	var nilai = 10
	switch {
	case nilai == 10:
		fmt.Println("Sempurna")
		fallthrough
	case (nilai < 10) && (nilai >= 7):
		fmt.Println("Bagus")
	case (nilai <= 6) && (nilai >= 4):
		fmt.Println("Lumayan")
	default:
		fmt.Println("Kurang")
		fmt.Println("Coba lagi")
	}
}
