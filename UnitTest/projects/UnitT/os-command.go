package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	arg1 := args[0] // 첫번째 인자 프로그램 이름 반환
	fmt.Println("args[0] prefix : ", arg1)

	restArgs := args[1:]
	fmt.Println(restArgs)

	for idx, arg := range restArgs {
		fmt.Printf("%d번째 인자: %s\n", idx, arg)
	}
}

/*
	go run .\os-command.go "fluxtest" "188"

	go build os-command.go

	PS D:\Works\Go\projects\UnitT> .\os-command "fluxtest" "188"

*/
