package main

import (
	"fmt"
)

var a int = 100
var b int = a
var arrayA []int = []int{1, 2, 4, 5}
var mapA map[int]string = map[int]string{}
var c *int = &a

func main() {
	example1()
	fmt.Println()
	example2()
	fmt.Println()
	example3()
	fmt.Println()
	example4()
	fmt.Println()
	example5()
}

func example1() {
	// Print
	fmt.Println("Harga original a adalah: ", a)
	fmt.Println("Harga original b adalah: ", b)

	// Perubahan 1 (Perubahan pada normalnya)
	b = 1000

	fmt.Println("Harga a setelah perubahan 1 adalah: ", a)
	fmt.Println("Harga b setelah perubahan 1 adalah: ", b)
}

func example2() {
	// Perubahan 2 (Perubahan passed by value)
	changeWithoutPointer(a)
	changeWithoutPointer(b)

	fmt.Println("Harga a setelah perubahan 2 adalah: ", a)
	fmt.Println("Harga b setelah perubahan 2 adalah: ", b)
}

func example3() {
	// Perubahan 3 (Contoh Slice)
	var arrayA []int = []int{1, 2, 4, 5}
	fmt.Println("Harga anggoata ke-1 dari arrayA adalah: ", arrayA[0])
	changeArray(arrayA)
	fmt.Println("Harga anggoata ke-1 dari arrayA setelah diakukan perubahan adalah: ", arrayA[0])
}

func example4() {
	// Perubahan 4 (Contoh Map)
	var mapA map[int]string = map[int]string{}

	mapA[0] = "Hello"
	mapA[1] = "World"
	fmt.Println("Harga anggota ke-1 dari mapA adalah: ", mapA[0])
	changeMap(mapA)
	fmt.Println("Harga anggota ke-1 setelah dilakukan perubahan adalah", mapA[0])
}

func example5() {
	// Perubahan 5 (Contoh Menggunakan Pointer)
	var c *int = &a
	fmt.Println("Alamat dari c adalah: ", c)
	fmt.Println("Harga dari c adalah: ", *c)
	fmt.Println("__________________________")
	fmt.Println("Alamat dari a adalah: ", &a)
	fmt.Println("Harga dari a adalah: ", a)

	// Perubahan
	fmt.Println("__________________________")
	changeWithPointer(c)
	fmt.Println("Harga a setelah perubahan: ", a)
}

// Helpers
func changeWithoutPointer(data int) {
	data = -100
}

func changeWithPointer(data *int) {
	*data = 1000
}

func changeArray(data []int) {
	data[0] = -100
}

func changeMap(data map[int]string) {
	data[0] = "Ini bukan hello world"
}

func arrayExample() {
	var arrayA []int = []int{1, 2, 3}

	arrayA[0] = 1000

	var arrayB []string = []string{"a", "b"}

	arrayB[2] = "Luki"
}
