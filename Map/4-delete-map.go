package main

import "fmt"

func main() {
	var ayam = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    30,
		"april":    20,
	}

	fmt.Println(len(ayam))
	fmt.Println(ayam)

	delete(ayam, "januari")

	fmt.Println(len(ayam))
	fmt.Println(ayam)

}
