package main

import "fmt"

func main() {
	// konsep nya sama seperti interface yang ada di java
	// untuk sebuah kontrak kode

	// berbeda dengan bahasa lainnya yang secara eskplicit menyebutkan interface mana yang menjadi kontrak.
	// Kalo di go itu tidak perlu, yang penting kode nya implement method yang ada di dalam interface

	ayam := Animal{"ayam"}
	sayHello(ayam)

	// run interface kosong
	sayHello2()
	sayHello3([]string{"abdu", "eunha", "sowon"})
	sayHello3(true)
	sayHello4(false)
	sayHello4(ayam)
}

type Customer interface {
	GetName() string
	// Return() int
	// Walk(destination string)
}

// pakai interface
func sayHello(customer Customer) {
	fmt.Println("Hello, ", customer.GetName())
}

// contoh implementasi interface dengan struct
type Animal struct {
	Name string
}

// override method yang ada di interface
func (animal Animal) GetName() string {
	return animal.Name
}

// interface kosong (tidak ada method yang dideklarasikan)
// interface kosong itu merupakan interface yang di implementasi kan oleh semua tipe data di Go
// sehingga jika kita set interface{} maka kita bisa isi didalam nya dengan tipe data apapun
// alias type nya adalah `any`
func sayHello2() interface{} {
	// return 1
	// return true
	return "abdu"
}

func sayHello3(elemen ...any) {
	for _, element := range elemen {
		fmt.Println("Hello, ", element)
	}
}

func sayHello4(nama interface{}) {
	fmt.Println(nama)
}
