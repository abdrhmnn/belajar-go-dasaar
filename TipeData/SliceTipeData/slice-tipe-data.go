package main

import "fmt"

func main() {
	// tipe data Slice, yaitu tipe yang berisikan sebagian data array
	// Slice memiliki 3 konsep penting -> pointer, length dan capacity
	name := [...]string{"abdu", "eunha", "1", "2", "200", "5000"}

	// slice dri index 0 sampai 3
	fmt.Println(name[0:4])
	// pointer -> 0
	// length -> 4
	// capacity -> 6

	fmt.Println(name[:3]) // slice array dari index pertama sampai index 2
	fmt.Println(name[3:]) // slice array dari index 3 sampai terakhir
	fmt.Println(name[:])  // slice array dari index awal sampai akhir

	// kalo mau deklarasi menjadi tipe data slice, maka
	// var sliceArray []string = name[:]
	// sliceArray2 := name[:]

	// functional slice
	fmt.Println(len(name[:2])) // length slice
	fmt.Println(cap(name[5:])) // capacity slice
	newSlice := name[2:4]
	newSlice[0] = "sowon"

	newSlice2 := append(newSlice, "20", "30") // make new slice with append with old slice

	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	newSlice3 := make([]int16, 2, 10) // make new slice and array together
	newSlice3[0] = 100
	newSlice3[1] = 100
	newSlice3[10] = 200
	fmt.Println(len(newSlice3))

	fromSlice := name[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice) // copy slice

	// perbedaan deklarasi array dan slice
	iniArray := [...]bool{true, false, false}
	iniSlice := []int32{100, 200, 20, 30, 45}
}
