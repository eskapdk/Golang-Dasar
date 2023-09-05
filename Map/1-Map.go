package main

import "fmt"

func main() {
	var ayam map[string]int
	ayam = map[string]int{}

	ayam["Januari"] = 50
	ayam["Februari"] = 40

	fmt.Println("Januari", ayam["Januari"])
	fmt.Println("Februari", ayam["Februari"])
	fmt.Println("Maret", ayam["Maret"])
}
