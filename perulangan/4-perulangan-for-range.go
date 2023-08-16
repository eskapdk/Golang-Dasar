package main

import "fmt"

func main() {
	var nilai1 = "12345"
	fmt.Println("=======STRING=======")
	for i, v := range nilai1 {
		fmt.Println("Index : ", i, "Nilai : ", v)
	}

	var nilai2 = [5]int{11, 22, 33, 44, 55}
	fmt.Println("=======ARRAY=======")
	for _, v := range nilai2 {
		fmt.Println("Nilai : ", v)
	}

	var nilai3 = nilai2[1:4]
	fmt.Println("=======SLICE=======")
	for _, v := range nilai3 {
		fmt.Println("Nilai : ", v)
	}

	var nilai4 = map[byte]int{'A': 23, 'B': 34, 'C': 45}
	fmt.Println("=======MAP=======")
	for k, v := range nilai4 {
		fmt.Println("index :", k, "NIlai :", v)
	}

	for range nilai4 {
		fmt.Println("Done")
	}
}
