package main

import "fmt"

func main() {
	// map adalah tipe data `pair` (key-value)
	// bisa dimasukkan sebanyak-banyaknya

	nama := map[int32]string{
		10: "abdu",
		20: "sowon",
	}
	nama[200] = "jkt"
	nama[1_000_000] = "bdg"

	fmt.Println(nama)
	fmt.Println(nama[200])

	// functional map
	fmt.Println(len(nama))
	fmt.Println(nama[100])

	delete(nama, 20)                 // delete
	newNama := make(map[string]bool) // make new map without values
}
