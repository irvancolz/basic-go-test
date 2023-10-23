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

	// baca catatan
	copyslice := sliceOfNumber[20] == 99
	fmt.Println(sliceOfNumber, copyslice)

	// catatan

	// - slice dan array menggunakan indexing sebagai caa mengidentifikasi setiap value di dalamnya
	// dan indexing di mulai dari angka 0, oleh karena itu dalam mengakases data terakhir dari sebuah array atau slice dengan panjang n tidak
	// dapat menggunakan cara array[n] namun harus diperhatikan lagi karena indexing dimulai dari angka 0 sehingga pengkasesan elemen terakhir
	// yang paling umum digunakan adalah dengan array[n - 1]

	// - perbedaan operator '=' dan '==', operator '='digunakan untuk mengassign sebuah value ke variable, sedangkan operator
	// '==' disebut sebagai operator perbandingan, operator ini digunakan apakah dua buah value/variable yang di bandingkan memiliki
	// nilai yang sama atau tidak
}
