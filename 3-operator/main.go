package main

import "fmt"

func main() {
	var a = 10
	var b = 5
	// Arithmetic
	fmt.Println(a + b) // 10 + 5
	fmt.Println(a - b) // 10 - 5
	fmt.Println(a * b) // 10 * 5
	fmt.Println(a / b) // 10 / 5
	fmt.Println(a % b) // 10 % 5
	fmt.Println("")
	// Assignment
	c := a // c = 10
	d := b // d = 5
	c += d // 10 + 5
	fmt.Println(c)
	c -= d // 15 - 5
	fmt.Println(c)
	c *= d // 10 * 5
	fmt.Println(c)
	c /= d // 50 / 5
	fmt.Println(c)
	c %= d // 10 % 5
	fmt.Println(c)
	fmt.Println("")
	// Logical
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a != b)
	fmt.Println(a == b)
}
