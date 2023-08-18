package main

import "fmt"

func main() {
	var angka1 = [2][3]int{[3]int{1, 2, 3}, [3]int{1, 2, 3}}
	var angka2 = [2][3]int{{3, 2, 1}, {3, 2, 1}}

	fmt.Println("Angka 1 :", angka1)
	fmt.Println("Angka 2 :", angka2)
}
