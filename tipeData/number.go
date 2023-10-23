package tipedata

import "fmt"

func Number() {

	//1. buatlah sebuah variable dengan nama 'Age' dengan tipe data Int8
	//2. buatlah sebuah variable dengan nama 'Balance' dengan tipe data Int64
	//3. buatlah sebuah variable  jokoAge dengan tipe data Int lalu inisiasikan nilainya dengan 32,
	// 	buat variable susiAge dengan tipe data int lalu inisialisasikan nilainya dengan 24, setelah itu lakukan operasi
	//  pengurangan terhadap variable jokoAge dengan susiAge, lalu simpan nilainnya kedalam variable jokoAndSusiAgeGap

	// 4. budi memiliki usaha pembuatan layang-layang, setiap layang layang yang budi buat membutuhkan batang bambu sejumlah 5,
	// hari ini budi berencana untuk membuat layang - layang dengan jumlah batang bambu 43 buah, tuliskan kode untuk mencari tahu berapa
	// jumlah sisa batang bambu yang tersisa setelah proses pembuatan layang-layang setiap hari (jumlah batang bambu dapat berubah setiap harinya,
	//  namun batang bambu yang dibutuhkan tetap 5 batang per layang-layang)

	// var Age int8
	// var Balance int64 // ini untuk jawaban no 1 dan 2
	jokoAge := 32
	susiAge := 24

	jokoAndSusiAgeGap := jokoAge - susiAge
	fmt.Println(jokoAndSusiAgeGap)

	batangbambuperlayang := 5
	batangbambuharini := 43

	// baca catatan
	sisabatangbambusetelahpembuatanuntukhariini := batangbambuharini / batangbambuperlayang
	fmt.Println(sisabatangbambusetelahpembuatanuntukhariini)

	// catatan
	// untuk mencari sisa hasil bagi dapat menggunakan operator %, berbeda dengan operator / yang menghasilkan
	// exact hasil pembagian contoh 7/3 = 2.3333, sedangkan 7%3 = 1.
}
