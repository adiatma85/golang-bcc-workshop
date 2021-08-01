package main

import "fmt"

func main() {
	fmt.Println("Saya berjanji tidak akan terlambat lagi")
	fmt.Println("")
	// for i := 0; i < 100; i++ {
	// 	fmt.Println("Saya berjanji tidak akan terlambat lagi")
	// }
	fmt.Println("")
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i+1, "Saya berjanji tidak akan terlambat lagi")
	// }
	// ganjil := 1

	// for ganjil < 10 {
	// 	fmt.Println(ganjil)
	// 	ganjil += 2
	// }

	genap := 0

	for {
		if genap%2 != 0 {
			continue
		}

		if genap >= 15 {
			break
		}

		fmt.Println(genap)

		genap += 2
	}
}
