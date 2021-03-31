package main

import (
	"flag"
	"fmt"
	"strings"
)

type StringArray []string

func (arr *StringArray) String() string {
	return fmt.Sprintf("%v", *arr)
}

func (arr *StringArray) Set(s string) error {
	*arr = strings.Split(s, ",")
	return nil
}

func main() {
	maxValue := flag.Int64("max", 10, "Defines maximum value")

	var minValue int64
	flag.Int64Var(&minValue, "min", 0, "Defines minimum value")

	var arrStrs StringArray
	flag.Var(&arrStrs, "array", "An inputted array that can be iterated")

	flag.Parse()

	for minValue < *maxValue {
		fmt.Printf("%v\n", arrStrs)
		minValue++
	}
}

/*
	go run .\flag-command.go -max=20 -min=15 -array=HelloWorld!

	go build flag-command.go

	PS D:\Works\Go\projects\UnitT> .\flag-command -max=20 -min=15 -array=HelloWorld!

*/
