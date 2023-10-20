package tipedata

import "fmt"

func main() {
	// 1. buat variable listOfNumber dengan tipe data array of int dengan panjang 10
	// isi datanya dengan nilai dari 1 -10
	//  rubah data dengan index ke 3 dengan nilai 99
	listOfNumber := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	listOfNumber[3] = 99
	fmt.Println(listOfNumber)

	// 2. buat variable sliceOfNumber dengan tipe data slice of int
	// isi datanya dengan nilai dari 1 - 20
	//  rubah data terakhir dengan nilai 99
	sliceOfNumber := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	copyslice := sliceOfNumber[20] == 99
	fmt.Println(sliceOfNumber, copyslice)
}
