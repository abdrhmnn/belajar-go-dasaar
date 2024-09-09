package main

import (
	"fmt"
	"package-import/helper"
	_ "package-import/pkginit"
)

func main() {
	fmt.Println("testing")
	helper.SayHello()
	// helper.IsLogin = true
}
