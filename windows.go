/*
 * @Author: RyanWilson
 * @Date: 2022-11-05 16:59:32
 */
package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	kernel32DLL     = syscall.MustLoadDLL("User32.dll")
	procEnumWindows = kernel32DLL.MustFindProc("EnumWindows")
)

func StringToCharPtr(str string) *uint8 {
	chars := append([]byte(str), 0)
	return &chars[0]
}

// 回调函数，用于EnumWindows中的回调函数，第一个参数是hWnd，第二个是自定义穿的参数
func AddElementFunc(hWnd syscall.Handle, hWndList *[]syscall.Handle) uintptr {
	*hWndList = append(*hWndList, hWnd)
	return 1
}

// 获取桌面下的所有窗口句柄，包括没有Windows标题的或者是窗口的。
func GetDesktopWindowHWND() {
	var hWndList []syscall.Handle
	hL := &hWndList
	r1, _, err := syscall.Syscall(procEnumWindows.Addr(), 2, uintptr(syscall.NewCallback(AddElementFunc)), uintptr(unsafe.Pointer(hL)), 0)
	if err != 0 {
		fmt.Println(err)
	}
	fmt.Println(r1)
	fmt.Println(hWndList)
}
