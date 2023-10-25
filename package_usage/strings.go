package packageusage

import (
	"fmt"
	"strings"
)

// buatlah function untuk menyelesaikan permasalahan berikut
func merubahstring() {
	//  1. membuat text menjadi UPPERCASE
	fmt.Println(strings.ToUpper("text"))
	//  2. membuat text menjadi lowercase
	fmt.Println(strings.ToLower("TEXT"))
	//  3. merubah text seperti "lorem ipsun dolor sit amet" menjadi array berisi [lorem, ipsum, dolor, sit, amet]
	splittedArr := strings.Split("lorem ipsun dolor sit amet", "olo")
	fmt.Println("hasil =", splittedArr, " total element :", len(splittedArr))
	//  4. merubah text seperti "lorem&+ipsun&+dolor&+sit&+amet" menjadi array berisi [lorem, ipsum, dolor, sit, amet]
	fmt.Println(strings.Split("lorem&+ipsun&+dolor&+sit&+amet", "&+"))
	//  5. merubah array dengan format [lorem, ipsum, dolor, sit, amet] menjadi string seperti "lorem&+ipsum&+dolor&+sit&+amet"
	arrOfStr := []string{"lorem", "ipsum", "dolor", "sit", "amet"}
	fmt.Println(strings.Join(arrOfStr, "&+"))
	//  6. membandingkan dua buah string apakah sama memperdulikan huruf kapital / kecil dari kedua text

	texta := "text"
	textb := "Text"

	if texta == textb {
		fmt.Println("string sama")
	} else {
		fmt.Println("string tidak sama")
	}
	//  7. membandingkan dua buah string apakah sama tanpa memperdulikan huruf kapital / kecil dari kedua text
	text1 := "text"
	text2 := "Text" // text == Text
	if strings.EqualFold(text1, text2) {
		fmt.Println("string sama")
	} else {
		//  text != Text
		fmt.Println("string tidak sama")
	}
}
