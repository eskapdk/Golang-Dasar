package main

import "fmt"

func main() {
	var nilai = 8
	if nilai >= 7 {
		switch nilai {
		case 10:
			fmt.Println("Sempurna")
		default:
			fmt.Println("Kerja Bagus")
		}
	} else {
		if nilai == 5 {
			fmt.Println("Lumayan")
		} else if nilai == 3 {
			fmt.Println("Perlu Belajar Lagi")
		} else {
			fmt.Println("Ditingkatkan Belajarnya")
			if nilai == 0 {
				fmt.Println("Gagal")
			}
		}
	}
}
