package main

import (
	"fmt"
)

func main() {
	sayHello()
	sayHello2("abdu", 200)
	sayHello3("abdu", 200)

	// contoh jika ingin tangkap masing2 return ke dalam variabel
	resultInt, resultBool, resultString := sayHello4("abdu", 200)

	// ignore catching return value
	_, _, resultString2 := sayHello4("abdu", 200)

	a, b := sayHello5()

	sayHello6("abdu", 10, 20, 30, 50000, 2001021232, 202020)
	sayHello6("abdu", []int{10, 20, 30, 50000, 2001021232, 202020}...)

	// function as value
	funcValue := sayHello3
	fmt.Println(funcValue("abdu", 200))

	// function as parameter
	sayHello7("ABDU", spamFilter)

	// anonymous fucntion
	newAnonymoutFunc := func(string) string {
		return "eunha"
	}

	sayHello8("abdu", newAnonymoutFunc)
	sayHello8("abdu", func(string) string {
		return "eunha"
	})
}

func sayHello() {
	fmt.Println("test")
}

// void
func sayHello2(nama string, umur int) {
	fmt.Println(nama, umur)
}

// return value
func sayHello3(nama string, umur int) int {
	fmt.Println(nama, umur)
	return 200
}

// multiple return values
func sayHello4(nama string, umur int) (int, bool, string) {
	fmt.Println(nama, umur)
	return 200, true, "abdu"
}

// named return values
func sayHello5() (nama string, umur int) {
	nama = "abdu"
	umur = 2000
	return nama, umur
}

// variadic function
func sayHello6(nama string, umur ...int) {
	for _, hasil := range umur {
		fmt.Println(hasil)
	}
	fmt.Println(nama)
}

// function as parameter
func sayHello7(nama string, filter func(string) string) {
	fmt.Println("Hai ", nama, filter(nama))
}

func spamFilter(nama string) string {
	if nama == "duar" {
		return "..."
	} else {
		return nama
	}
}

// function as parameter with type declaration
type Filter func(string) string

func sayHello8(nama string, filter Filter) {
	fmt.Println("Hai ", nama, filter(nama))
}
