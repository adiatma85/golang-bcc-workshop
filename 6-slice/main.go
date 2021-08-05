package main

import "fmt"

func main() {
	// example1()
	// example2()
	example3()
}

// Array in Golang
func example1() {
	// Array adalah kumpulan data. Statis, sehingga tidak bisa ditambah (append)
	var ArrayA [2]int = [2]int{1, 2}
	fmt.Println(ArrayA)

	// Cara Assign array lain dengan menggunakan loop
	var ArrayB [3]int
	// Loop versi satu
	for i := 0; i < len(ArrayB); i++ {
		ArrayB[i] = i
	}
	// Loop versi dua
	for _, item := range ArrayB {
		fmt.Println(item)
	}

	// Atau secara manual
	var ArrayC [2]int
	ArrayC[0] = 1
	ArrayC[1] = 2

	// Array tidak dapat di append
	// Akan menunjukkan error Out Of Bond
	// ArrayC[2] = 3
}

// Slice in Golang
func example2() {
	// Slice adalah potongan dari Array, bukan Array
	var ArrayD [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Slice
	var SliceA []int = ArrayD[:]

	fmt.Println(SliceA)

	// Slice lebih fleksibel dibandingkan dengan Array
	SliceA = append(SliceA, 100)
	fmt.Println(SliceA)
}

// Map in Golang
func example3() {
	// Map adalah pemetaan berbentuk key => value
	MapA := map[int]string{2: "Februari"}
	MapA[1] = "Januari"

	fmt.Println(MapA)
	fmt.Println(MapA[1])
	fmt.Println(MapA[2])
}
