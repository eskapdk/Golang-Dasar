package main

import "fmt"

func main() {
	var nilai = 10
	switch nilai {
	case 10:
		fmt.Println("Sempurna")
	case 9, 8, 7:
		fmt.Println("Bagus")
	case 6, 5:
		fmt.Println("Lumayan")
	default:
		fmt.Println("Coba lagi")
	}
}
