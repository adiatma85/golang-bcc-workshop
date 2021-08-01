package main

import "fmt"

func main() {
	var boolean1 bool = true
	fmt.Println(boolean1)
	var sentence1 = "lorem"
	fmt.Println(sentence1)
	var sentence2 = 3
	fmt.Println(sentence2)
	var sentence3 = ""
	fmt.Println(sentence3)
	var number1 int = 123
	fmt.Println(number1)
	var number2 int = "ipsum"
	fmt.Println(number2)
	var number3 int = 123.3
	fmt.Println(number3)
	var number4 uint = -4
	fmt.Println(number4)
	var number5 float64 = 123.3
	fmt.Println(number5)
	var number6 float64 = 123
	fmt.Println(number6)
	var number7 float64 = 0.0
	fmt.Println(number7)
	var number8 float64 = nil
	fmt.Println(number8)

	var first, second, third string
	first, second, third = "ichi", "ni", "san"
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)
}
