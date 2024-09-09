package main

import (
	"errors"
	"fmt"
)

// untuk handle error dengan menggunakan interface error bawaan dari go-lang
// biasanya menggunakan helper `errors`, seperti cth dibawah

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Something goes wrong!")
	} else {
		return nilai / pembagi, nil
	}
}

// custom error dengan menggunakan interface `error` bawaan go-lang
type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "id kosong!"}
	}

	if id != "abdu" {
		return &notFoundError{Message: "data: abdu, tidak ditemukan!"}
	}

	return nil
}

func main() {
	hasil, err := pembagian(100, 10)
	if err == nil {
		fmt.Println("Hasil: ", hasil)
	} else {
		fmt.Println("Error: ", err.Error())
	}

	err2 := SaveData("", nil)
	if err2 != nil {
		if validationError, ok := err2.(*validationError); ok {
			fmt.Println("Validation error: ", validationError.Error())
		} else if notFoundError, ok := err2.(*notFoundError); ok {
			fmt.Println("Not Found error: ", notFoundError.Error())
		} else {
			fmt.Println("Unknown error")
		}
	} else {
		fmt.Println("No have error: success!")
	}
}
