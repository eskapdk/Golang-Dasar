package main

import "fmt"

func main() {

	var nilai = 8

	if nilai == 10 {
		fmt.Println("Lulus dengan nilai sempurna")
	} else if nilai > 5 {
		fmt.Println("Lulus")
	} else if nilai == 4 {
		fmt.Println("Hampir Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}
}
