package main

import "fmt"

func main() {
	// nil adalah data kosong
	// nil hanya bisa di digunakan di tipe data: interface, func, map, slice, pointer dan channel

	data1 := sayHello("abdu")
	fmt.Println(data1)
}

func sayHello(nama string) map[string]string {
	if nama == "" {
		return nil
	} else {
		return map[string]string{
			"nama": nama,
		}
	}
}
