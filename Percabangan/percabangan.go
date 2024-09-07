package main

import "fmt"

func main() {
	nama := "abdu"

	if nama == "abdu" {
		fmt.Println("oke")
	} else if nama == "eunha" {
		fmt.Println("test")
	} else {
		fmt.Println("okkk")
	}

	// if short statement
	if length := len(nama); length > 5 {
		fmt.Println("benar")
	} else {
		fmt.Println("salah")
	}

	// switch
	switch nama {
	case "abdu":
		fmt.Println("oke")
	case "eunha":
		fmt.Println("oke")
	default:
		fmt.Println("salah")
	}

	// switch short statement
	switch length := len(nama); length > 5 {
	case true:
		fmt.Println("oke")
	case false:
		fmt.Println("oke")
	}

	// switch tanpa kondisi
	length := len(nama)
	switch {
	case length > 5:
		fmt.Println("oke")
	case length == 10:
		fmt.Println("okkk")
	}
}
