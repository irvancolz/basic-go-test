package tipedata

import "fmt"

//1. buat struct Person dengan komposisi sbb :
// - memiliki property Name dengan tipe String
// - memiliki property Age dengan tipe int
// - memiliki property Is married dengan tipe bool
type Person struct {
	name    string
	age     int
	married bool
}

// 2. buat method sayHello untuk struct Person
// dalam method sayHello print "hello "+ nama person ke konsol
func (person Person) sayHello(name string) {
	fmt.Println("hello", person.name)
}

func Struct() {
	//  3. buat struct Person baru lalu isi datanya dengan data anda
	// lalu simpan datanya ke dalam variable baru
	var faris Person
	faris.name = "faris"
	faris.age = 21
	faris.married = false
	//  panggil method sayHello
	faris.sayHello("faris")
}
