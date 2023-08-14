package main

import "fmt"

func main() {

	// tanpa menggunakan var dan tipe data
	firstName := "Eska"
	// menggunakan var dan tanpa tipe data
	var lastName = "Pratama"

	fmt.Printf("Halo %s %s!\n", firstName, lastName)

	fmt.Printf("halo Eska Pratama!\n")
	fmt.Printf("halo %s!\n", firstName)
	fmt.Println("halo", firstName, lastName+"!")
}
