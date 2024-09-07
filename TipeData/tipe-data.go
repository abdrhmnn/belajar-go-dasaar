package main

import "fmt"

func main() {
	// tipe data number:
	// integer (bilangan bulat) -> int8, int16, int32, int64
	// integer (khusus bilangan positif 0 - sekian) -> uint8, uint16, uint32, uint64

	// tipe data floating point:
	// float32, float64, complex64, complex128

	// alias name untuk tipe data integer dan floating point :
	// byte -> int8
	// rune -> int32
	// int -> minimal int32 ke atas
	// uint -> minimal uint32 ke atas

	fmt.Println("Hasil: ", 1)
	fmt.Println("Hasil 2: ", 10.5)

	// boolean
	// true, false

	// string
	fmt.Println(len("abdu")) // length string
	fmt.Println("abdu"[2])   // ambil bagian dari string berdasarkan index

	// konversi tipe data
	var hasil int16 = 2_000
	var hasil2 int32 = int32(hasil)
	var hasil3 int8 = int8(hasil2)

	// alias type declaration
	type NoKTP string
	var ktpAbdu NoKTP = "22112121121"
	var contoh string = "111111"
	var konversiKtpAbdu NoKTP = NoKTP(contoh) // konversi tipe data lain ke alias type

	// array
	// tidak bisa mengubah length array setelah dibuat
	var arrNama [5]string
	arrNama[0] = "abdu"
	arrNama[1] = "abdu2"

	var arrAngka = [3]int32{10, 200} // kalo kita gak isi elemen yg gak sesuai dgn length array nya
	// maka akan diisi nilai default sesuai tipe data nya, int = 0, string = "", dll

	fmt.Println(len(arrNama))
	fmt.Println(arrNama[2])

	// bisa juga assign length array tanpa set length nya
	var arrAngka2 = [...]int32{200, 300, 500, 200_000} // tpi ini hanya bisa ketika kita set langsung value array nya
}
