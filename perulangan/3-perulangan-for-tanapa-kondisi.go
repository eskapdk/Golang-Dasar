package main

import "fmt"

func main() {
	var nilai = 0
	for {
		fmt.Println("Angka :", nilai)
		nilai++
		if nilai == 5 {
			break
		}
	}
}
