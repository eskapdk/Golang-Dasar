package main

import "fmt"

func main() {
	var firstName string = "Eska"
	var lastName string
	lastName = "Pratama"

	fmt.Printf("Halo %s %s!\n", firstName, lastName)

	fmt.Printf("halo Eska Pratama!\n")
	fmt.Printf("halo %s!\n", firstName)
	fmt.Println("halo", firstName, lastName+"!")
}
