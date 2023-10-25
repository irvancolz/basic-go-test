package packageusage

import (
	"log"
	"sort"
)

// urutkan data berikut agar terurut dari terkecil ke yang paling besar / a-z
func urutkanangka() {
	var arrOfIntReversed = []int64{8, 4564, 8654, 84, 653132, 654, 8564, 321, 534, 8, 9, 534, 6864, 8354, 65, 1, 3, 45, 34, 62, 63, 12, 56, 45634, 67, 34, 8, 4, 345, 5854, 478, 4, 6, 7345, 67324, 564, 78, 546, 468445, 57, 645343455, 72, 351, 84, 321, 51, 8, 32, 154, 984, 321, 84, 51, 213, 51, 8321, 564, 8, 325, 186, 4798, 516, 8498, 31, 84, 6, 31, 21, 56, 498, 54, 6786, 5612, 54, 87987968, 46749, 5, 13268, 84, 54}
	var arrOfStrReversed = []string{"z", "y", "w", "v", "u", "t", "s", "r", "q", "p", "o", "n", "s", "m", "l", "k", "j", "i", "h", "g", "f", "e", "d", "c", "b", "a"}

	//  urutkan data berikut adar terurut dari paling besak ke yang paling kecil / z-a
	var arrOfInt = []int64{8, 4564, 8654, 84, 653132, 654, 8564, 321, 534, 8, 9, 534, 6864, 8354, 65, 1, 3, 45, 34, 62, 63, 12, 56, 45634, 67, 34, 8, 4, 345, 5854, 478, 4, 6, 7345, 67324, 564, 78, 546, 468445, 57, 645343455, 72, 351, 84, 321, 51, 8, 32, 154, 984, 321, 84, 51, 213, 51, 8321, 564, 8, 325, 186, 4798, 516, 8498, 31, 84, 6, 31, 21, 56, 498, 54, 6786, 5612, 54, 87987968, 46749, 5, 13268, 84, 54}
	var arrOfStr = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "s", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	// untuk mengurutkan variabel arrOfInt besar ke kecil
	sort.Slice(arrOfInt, func(i, j int) bool {
		return arrOfInt[i] > arrOfInt[j]
	})
	log.Println(arrOfInt)

	// untuk mengurutkan variabel arrOfStr besar ke kecil
	sort.Slice(arrOfStr, func(i, j int) bool {
		return arrOfStr[i] > arrOfStr[j]
	})
	log.Println(arrOfStr)

	// untuk mengurutkan variabel arrOfIntReversed kecil ke besar
	sort.Slice(arrOfIntReversed, func(i, j int) bool {
		return arrOfIntReversed[i] < arrOfIntReversed[j]
	})
	log.Println(arrOfIntReversed)

	// untuk mengurutkan variabel arrOfStrReversed kecil ke besar
	sort.Slice(arrOfStrReversed, func(i, j int) bool {
		return arrOfStrReversed[i] < arrOfStrReversed[j]
	})
	log.Println(arrOfStrReversed)

}
