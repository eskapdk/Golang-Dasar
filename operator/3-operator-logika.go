package main

import "fmt"

func main() {
	var left = true
	var right = false

	fmt.Println("Nilai left :", left)
	fmt.Println("Nilai right :", right)

	var leftAndright = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndright)

	var leftOrright = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrright)

	var reverseRight = !right
	fmt.Printf("!right \t\t(%t) \n", reverseRight)
}
