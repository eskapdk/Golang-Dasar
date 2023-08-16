package main

import "fmt"

func main() {

	for i := 0; i <= 8; i++ {
		if i%2 != 1 {
			continue
		}
		if 1 > 8 {
			break
		}
		fmt.Println("Angka : ", i)
	}
}
