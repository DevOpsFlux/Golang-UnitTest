package main

import (
	"C"
	"fmt"
)
import "syscall"

/*
var (
	myWin32DLL            = syscall.NewLazyDLL("MyWin32DLL.dll")
	prncWin32CallbackTest = myWin32DLL.NewProc("Win32CallbackTest")
)

func myCallback() uintptr {
    fmt.Println("myCallback is Called!!!")
    return 0
}
*/
func main() {

	moduleKernel32 := syscall.NewLazyDLL("kernel32.dll")
	hModule := moduleKernel32.Handle() // base address of kernel32.dll at runtime

	fmt.Println("myCallback is Called!!!")

	/*
		prncWin32CallbackTest.Call(
			uintptr(syscall.NewCallback(myCallback))
		)
	*/
}
