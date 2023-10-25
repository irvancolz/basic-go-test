package packageusage

import (
	"fmt"
	"time"
)

func waktu() {
	//  1. buatlah tipe data waktu baru dengan tanggal 24 desember 2022 pukul 12 malam ke dalam variable shibuyaAccident
	shibuyaAccident := time.Date(2022, time.December, 24, 0, 0, 0, 0, time.UTC)
	fmt.Println(shibuyaAccident)
	//  2. ambil tanggal dari variabel shibuyaAccident ke dalam variable baru
	variabelbaru1 := shibuyaAccident.Format("24 december 2022")
	fmt.Println(variabelbaru1)
	//  3. ambil waktu dari variabel shibuyaAccident ke dalam variable baru
	variabelbaru2 := shibuyaAccident.Format("00:00;00")
	fmt.Println(variabelbaru2)
	//  4. format shibuya Accident menjadi 24 desember 2022 pukul 12 malam dan simpan hasilnya dalam variable baru
	variabelbaru3 := shibuyaAccident.Format("24 desember 2022 pukul 12 malam")
	fmt.Println(variabelbaru3)
	//  5. rubah waktu shibuyaAccident menjadi format *EPOCH second
	variabelbaru4 := shibuyaAccident.Unix()
	fmt.Println(variabelbaru4)
	//  6. buat waktu ini menjadi format waktu time.Time = 1698150891
	waktu := 1698150891
	waktutime := time.Unix(int64(waktu), 0)
	fmt.Println(waktutime)
}
