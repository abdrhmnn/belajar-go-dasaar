package main

import "fmt"

func main() {
	// adalah template data yg digunakan untuk menggabungkan 0 (nol) atau lebih tipe data lainnya dalam
	// satu kesatuan, data di dalam struct itu disimpan di dalam field atau bisa dikatakan
	// struct adalah kumpulan dari field

	// cara menggunakan struct
	var customer1 Customer // initialisasi struct
	customer1.Nama = "abdu"
	fmt.Println(customer1.Nama)

	// cara lain
	customer2 := Customer{
		Nama:    "eunha",
		Address: "jkt",
		Umur:    200,
	}
	customer2.sayHello()

	customer3 := Customer{"sowon", "bdg", 300}
	customer3.sayHello()
}

type Customer struct {
	Nama    string
	Address string
	Umur    int
}

// struct method
func (customer Customer) sayHello() {
	fmt.Println("Hello, ", customer.Nama)
}
