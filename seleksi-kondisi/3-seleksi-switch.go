package main

import "fmt"

func main() {
	var nilai = 9

	switch nilai {
	case 8:
		fmt.Println("Sempurna")
	case 7:
		fmt.Println("Bagus")
	default:
		fmt.Println("Coba Lagi")
	}
}
