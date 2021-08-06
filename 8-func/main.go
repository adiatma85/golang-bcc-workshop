package main

import "fmt"

func main() {

	// fmt.Println("Hello World")

	fmt.Println(example1())

	// fmt.Println(example2())

	// Slice
	// var function = func(variable int) int {
	// 	return variable
	// }
	// example3(function, 4)
}

func example1() string {
	return "Hello World"
}

func example2() (iniString string) {
	iniString = "Heheh"
	return iniString
}

func example3(iniFungsi func(a int) int, originalA int) int {
	return iniFungsi(originalA)
}

// Contoh lainnya
func findMax(numbers []int, max int) (int, []int) {
	var hasil []int
	for _, number := range numbers {
		if number < max {
			hasil = append(hasil, number)
		}
		continue
	}

	return max, numbers
}
