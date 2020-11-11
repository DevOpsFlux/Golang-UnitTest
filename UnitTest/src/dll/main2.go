package main

import (
	"fmt"
	"os"
	"syscall"
)

// ansi color
const (
	black  = "\033[1;30m%s\033[0m"
	red    = "\033[1;31m%s\033[0m"
	green  = "\033[1;32m%s\033[0m"
	yellow = "\033[1;33m%s\033[0m"
	blue   = "\033[1;34m%s\033[0m"
	purple = "\033[0;36m%s\033[0m"
	cyan   = "\033[0;36m%s\033[0m"
	white  = "\033[0;37m%s\033[0m"
)

func init() {
	// cmd 에서 ansi color 활성화 시키는 부분
	handle := syscall.Handle(os.Stdout.Fd())
	kernel32DLL := syscall.NewLazyDLL("kernel32.dll")
	setConsoleModeProc := kernel32DLL.NewProc("SetConsoleMode")
	setConsoleModeProc.Call(uintptr(handle), 0x0001|0x0002|0x0004)
}

func main() {
	fmt.Printf(black, "black")
	fmt.Println("")
	fmt.Printf(red, "red")
	fmt.Println("")
	fmt.Printf(green, "green")
	fmt.Println("")
	fmt.Printf(yellow, "yellow")
	fmt.Println("")
	fmt.Printf(blue, "blue")
	fmt.Println("")
	fmt.Printf(purple, "purple")
	fmt.Println("")
	fmt.Printf(cyan, "cyan")
	fmt.Println("")
	fmt.Printf(white, "white")
	fmt.Println("")
}
