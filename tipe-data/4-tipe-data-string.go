package main

import "fmt"

func main() {
	//tipe data string biasa
	var message string = "Halo"
	fmt.Printf("message: %s \n", message)
	//tipe data string dengan backticks (``)
	var message2 = `Nama saya "Eska Pratama".
Salam kenal.
Mari belajar "Golang".`
	fmt.Println(message)
	fmt.Println(message2)
}
