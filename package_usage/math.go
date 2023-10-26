package packageusage

import (
	"fmt"
	"math"
)

// tulislah kode program yang dapat menyelesaikan permasalahan berikut

// Andi adalah seorang pemilik toko penyewaan bola. Dia memiliki 63 bola sepak, 48 bola basket, dan 29 bola voli
// yang bisa disewakan kepada pelanggan. Setiap kali seorang pelanggan ingin menyewa bola,
// dia akan memberikan 5 bola per jenisnya (untuk sepak bola, bola basket, dan bola voli).
// Namun, Andi ingin memastikan bahwa setiap kali dia memberikan bola, jumlah bola yang tersisa per jenisnya adalah setidaknya 10.
// 1. Berapa total paket yang dapat disewakan oleh Andi sebelum dia harus menyediakan lebih banyak bola?

func PenyewaanBola() {

	// - aturan penyewaan
	//  1. setiap kali menyewa akan di berikan 5 bola tiap jenisnya (1 paket = 5 voli, 5 basket, 5 sepak)
	//  2. setidaknya harus ada 10 bola dari setiap jenis yang tersisa

	bolasepak := 63
	bolabasket := 48
	bolavoli := 29
	paketpenyewaan := 5
	jumlahyangharustersisaperjenisnya := 10
	totalbolayangdapatdisewakan := 0

	for bolasepak >= jumlahyangharustersisaperjenisnya && bolabasket >= jumlahyangharustersisaperjenisnya && bolavoli >= jumlahyangharustersisaperjenisnya {
		bolasepak -= paketpenyewaan
		bolabasket -= paketpenyewaan
		bolavoli -= paketpenyewaan

		totalbolayangdapatdisewakan += paketpenyewaan
	}

	fmt.Println("Total bola yang dapat disewakan sebelum harus menyediakan lebih banyak:", totalbolayangdapatdisewakan)
}

// Rina pergi berbelanja ke supermarket untuk membeli barang-barang rumah tangga. Dia memiliki daftar belanja yang mencantumkan harga dan jumlah barang yang harus dibeli.
// Berikut adalah beberapa barang yang ada di daftar belanja Rina:
// - Deterjen seharga Rp 27.500 per botol, dan dia perlu membeli 3 botol.
// - Gula seharga Rp 9.750 per kilogram, dan dia perlu membeli 2,5 kilogram.
// - Minyak goreng seharga Rp 18.250 per liter, dan dia perlu membeli 4 liter.
// - Sabun mandi seharga Rp 12.750 per bungkus, dan dia perlu membeli 5 bungkus.
// - Rina ingin mengetahui perkiraan total belanjaannya, tetapi dia tahu bahwa harga dan jumlah barang pada daftar belanjanya mungkin memiliki pecahan desimal. Bantu Rina menghitung perkiraan total belanjannya.

// Hitung total belanjaan Rina sebelum pembulatan.
// Hitung total belanjaan Rina setelah pembulatan ke angka terdekat.

func rinabelanja() {
	deterjenperbotol := 27500
	gulaperkilogram := 9750
	minyakgorengperliter := 18250
	sabunmandiperbungkus := 12750
	deterjenyangharusdibeli := 3
	gulayangharusdibeli := 2.5
	minyakgorengyangharusdibeli := 4
	sabunmandiyangharusdibeli := 5
	totalbelanja :=
		(deterjenperbotol * deterjenyangharusdibeli) +
			(gulaperkilogram * int(gulayangharusdibeli)) +
			(minyakgorengperliter * minyakgorengyangharusdibeli) +
			(sabunmandiperbungkus * sabunmandiyangharusdibeli)

	// ini untuk menampilkan sebelum pembulatan
	fmt.Println(totalbelanja)
	// ini untuk menampilkan setelah pembulatan ke angka terdekat
	fmt.Println(math.Round(float64(totalbelanja)))
}
