package main

import "fmt"

func main() {
	score := 55
	if score > 80 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B+")
	} else if score > 69 {
		fmt.Println("B")
	} else if score > 60 {
		fmt.Println("C+")
	} else if score > 55 {
		fmt.Println("C")
	} else if score > 50 {
		fmt.Println("D+")
	} else if score > 44 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}
