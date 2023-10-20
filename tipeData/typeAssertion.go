package tipedata

import "fmt"

func typeAssertion() {
	// 1. rubah tipe data dari variable num ke tipe data int64 ,int8, float64 dan simpan maisng masing hasilnya kedalam variable sendiri
	num := 10
	var int64Num = int64(num)
	var int8Num = int8(num)
	var float64Num = float64(num)

	fmt.Println("num", num)
	fmt.Println("int64Num", int64Num)
	fmt.Println("int8Num", int8Num)
	fmt.Println("float64Num", float64Num)

	//  2. rubah tipe data nama menjadi string
	var nama interface{}
	nama = "John Doe"
	namaString := nama.(string)
	fmt.Println(namaString)

	// 3. rubah tipe data menjadi []byte lalu rubah lagi menjadi string
	var alamat string
	alamat = "Jl Lorem Ipsum"
	alamatBytes := []byte(alamat)
	alamatString := string(alamatBytes)
	fmt.Println(alamatString)

}
