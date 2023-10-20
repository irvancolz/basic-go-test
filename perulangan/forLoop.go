package perulangan

import "fmt"

func perulanganFor() {
	//  1. buat perulangan sebanyak 10 kali dengan setiap perulangannya melakukan print ke konsol berupa "ini merupakan perulangan ke (jumlah perulangan)"
	for i := 0; i < 10; i++ {
		fmt.Println("ini merupakan jumlah perulangan ke", i)
	}
	//  2. buat perulangan sebanyak 10 kali dengan setiap perulangannya melakukan print ke konsol berupa "ini merupakan perulangan ke (jumlah perulangan)"
	// 	apabila mencapai perulangan ke 3 hentikan perulangan tersebut
	for i := 0; i < 10; i++ {
		fmt.Println("ini merupakan jumlah perulangan ke", i)
		if i == 3 {
			break
		}
	}
	//  2. buat perulangan sebanyak 10 kali dengan setiap perulangannya melakukan print ke konsol berupa "ini merupakan perulangan ke (jumlah perulangan)"
	// 	apabila mencapai perulangan yang merupakan kelipatan 3 skip perulangan tersebut dan langsung ke perulangan berikutnya
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Println("ini merupakan jumlah perulangan ke", i)
	}
	//  2. buat perulangan sebanyak 30 kali dengan setiap perulangannya melakukan print ke konsol berupa "ini merupakan perulangan ke (jumlah perulangan)"
	// 	apabila mencapai perulangan yang merupakan kelipatan 7 print "ini adalah kelipatan 7"
	for i := 0; i < 30; i++ {
		if i == 7 {
			fmt.Println("ini adalah kelipatan 7")
		}
		fmt.Println("ini merupakan jumlah perulangan ke", i)
	}
}

func main() {
}
