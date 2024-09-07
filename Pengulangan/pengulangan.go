package main

import "fmt"

func main() {
	counter := 0

	for counter < 10 {
		fmt.Println("ok", counter)
		counter++
	}

	for abdu := 0; abdu < 10; abdu++ {
		fmt.Println("ok", abdu)
		counter++
	}

	// loop data collection (array, slice, map), for range
	names := []string{"abdu", "eunha", "sowon"}
	for index, name := range names {
		fmt.Println("index ke-", index, " dan hasil: ", name)
	}

	for abdu := 0; abdu < len(names); abdu++ {
		fmt.Println("ok", abdu)
		counter++
	}

	for _, name := range names {
		fmt.Println("elemen ke- ", name)
	}

	// break and continue
	// break untuk stop perulangan
	// continue untuk lanjut ke next loop
}
