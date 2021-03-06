package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	getPhysicallyInstalledSystemMemoryProc := kernel32.NewProc("GetPhysicallyInstalledSystemMemory")

	var memory uint64

	ret, _, err := getPhysicallyInstalledSystemMemoryProc.Call(uintptr(unsafe.Pointer(&memory)))
	fmt.Printf("GetPhysicallyInstalledSystemMemory, return: %+v\n", ret)
	fmt.Printf("GetPhysicallyInstalledSystemMemory, memory: %d\n", memory)
	fmt.Printf("GetPhysicallyInstalledSystemMemory, err: %s\n", err)

}
