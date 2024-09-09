package helper

import "fmt"

// ketika ingin membuat function dan function trsbt bisa di akses di luar package nya itu harus berikan penamaan
// PascalCase

// karena ini berhubungan dengan access modifier di go-lang
// setiap variable, type, struct, interface, func, dll yg nama nya di awali huruf besar maka semua itu bisa
// di akses di luar package nya

var isLogin = false // hanya bisa di akses di dalam package nya
var IsLogin = false // bisa di akses di luar package

type NamaLengkap string
type namaLengkap string

type Customer struct {
	Nama string
}

type customer struct {
	Nama string
}

func SayHello() {
	fmt.Println("test")
}

func sayHello2() {
	fmt.Println("test")
}
