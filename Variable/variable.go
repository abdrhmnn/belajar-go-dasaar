package main

import "fmt"

func main() {
	// setiap variable hanya bisa di assign satu tipe data

	var nama string
	nama = "abdu"
	fmt.Println(nama)

	var umur = 20
	var umur2 int8 = 100
	fmt.Println(umur)
	fmt.Println(umur2)

	// bisa juga deklarasi variable tanpa keyword var
	alamat := "jkt" // ini untuk deklarasi awal saja
	fmt.Println(alamat)

	// bisa membuat group variable
	var (
		nama2   = "abdu"
		alamat2 = "bdg"
	)

	// constan var
	const nama3 = "eunha"
	const isLogin bool = false

	const (
		nama4       = "umji"
		harga int64 = 200_000_000
	)
}
