package pkginit

import "fmt"

// bisa menambahkan sebuah function yang akan dijalankan ketika sebuah package itu di import/pakai
// cukup bikin func dengan nama `init`

func init() {
	fmt.Println("init dari package pkginit berjalan!")
}
