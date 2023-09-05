package main

import "fmt"

func main() {
	var ayam = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    30,
		"april":    20,
	}

	var value, isEksist = ayam["april"]

	if isEksist {
		fmt.Println(value)
	} else {
		fmt.Println("Item tidak ada")
	}
}
