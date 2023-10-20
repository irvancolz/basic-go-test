package function

import (
	"fmt"
)

// 1. buat function helloWorld yang akan mencetak "hello world" ke konsol
func helloWorld() {
	fmt.Println("Hello World")
}

//2. buat function sayHello yang akan menerima parameter berupa string, lalu cetak "hello " + nama ke konsol

func sayHello(name string) {
	fmt.Println("hello", name)
}

//3. buat function getTotal yyang akan menerima 3 buah parameter bertipe int, lalu kembalikan nilai dari (a x b) : c) * a + (a + b +c)

func getTotal(a int, b int, c int) {
	a = 10
	b = 10
	c = 10

	hasil := (a*b)/c*a + (a + b + c)
	fmt.Println(hasil)
}
