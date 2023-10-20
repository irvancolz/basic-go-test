package pengkondisian

import "fmt"

func pengkondisianIf() {
	const umur = 27
	var status string
	//  buat pengkondisian dengan ketentuan sebagai berikut
	// 1.  status=  anak-anak jika umur b <= 12 tahun
	if umur <= 12 {
		status = "anak-anak"
		fmt.Println(umur, status)
		// 1.  status=  remaja jika umur berada diantara 13-20 tahun
	} else if umur > 12 && umur <= 20 {
		status = "remaja"
		fmt.Println(umur, status)
		// 1.  status=  dewasa jika umur berada diantara 21-50 tahun
	} else if umur > 20 && umur <= 50 {
		status = "dewasa"
		fmt.Println(umur, status)
		// 1.  status=  manula jika umur  >= 51 tahun
	} else if umur >= 51 {
		status = "manula"
		fmt.Println(umur, status)
	} else {
		status = "ahli kubur"
		fmt.Println("saya sudah", status)
	}
}
