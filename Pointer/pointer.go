package main

import "fmt"

func main() {
	// by default, di go itu ketika ada variable yang dikirim ke function, method atau ke variable lain
	// sebenarnya itu yg dikirim adalah duplikasi dari value awal nya
	// maka dikatakan semua variable di passing by value bukan by reference

	// contoh
	customer1 := Customer{"abdu", "200", "jkt"}
	customer2 := customer1 // copy values from variable customer1

	customer2.Nama = "jkt"

	fmt.Println(customer1)
	fmt.Println(customer2)

	// jika ingin membuat passing by reference, maka disinilah pointer berfungsi
	// yaitu membuat reference ke lokasi data di memory yg sama tanpa menduplikasi data yg sudah ada
	customer3 := Customer{"eunha", "100", "bdg"}
	customer4 := &customer3              // pointer with & symbol
	var customer5 *Customer = &customer3 // versi manual nya

	customer4.Nama = "sowon"
	customer5.Nama = "yerin"

	fmt.Println("customer 3", customer3)
	fmt.Println("customer 4", customer4)
	fmt.Println("customer 5", customer5)

	// contoh membuat pointer berbeda dari pointer awal
	customer5 = &Customer{"umji", "300", "mlg"}
	fmt.Println("customer 5 with new pointer", customer5)

	// misal ada case semua variable itu harus mengacu ke pointer awal, gaboleh buat data untuk pointer baru
	// maka disini menggunakan `asterisk operator` with * symbol
	*customer4 = Customer{"duar", "500", "kalimantan"}
	fmt.Println("customer 3", customer3)
	fmt.Println("customer 4", customer4)

	// kalo pengen bikin pointer dengan nilai value awal adalah kosong maka bisa pakai keyword `new`
	customer6 := new(Customer)
	customer7 := customer6
	// var customer7 *Customer = customer6

	customer7.Nama = "Customer baru"
	fmt.Println("customer 6", customer6)
	fmt.Println("customer 7", customer7)

	// pointer as param value
	sayHello(&customer3)

	// pointer in method
	abdu := Abdu{"abdu"}
	abdu.sayHello2()
	fmt.Println(abdu.Nama)
}

type Customer struct {
	Nama, Umur, Alamat string
}

type Abdu struct {
	Nama string
}

// pointer as parameter function
func sayHello(customer *Customer) {
	customer.Nama = "eunha"
}

// pointer in method
func (abdu *Abdu) sayHello2() {
	abdu.Nama = "Mr. " + abdu.Nama
}
