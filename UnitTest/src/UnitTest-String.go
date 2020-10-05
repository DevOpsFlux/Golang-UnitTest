package main

import (
	"fmt"
	"stringutil"
)

func main() {
	test()
}

func test() {
	fmt.Println("string Reverse")
	val := stringutil.Reverse("ABCDEF")
	fmt.Println(val)
}
