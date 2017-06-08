package main

import (
	"log"
	"syscall"
	"unsafe"
)

var (
	user32, _     = syscall.LoadLibrary("user32.dll")
	messageBox, _ = syscall.GetProcAddress(user32, "MessageBoxW")
)

const (
/*
	MB_OK              = 0x00000000
	MB_OKCANCEL        = 0x00000001
	MB_YESNOCANCEL     = 0x00000003
	MB_ICONHAND        = 0x00000010
	MB_ICONQUESTION    = 0x00000020
	MB_ICONEXCLAMATION = 0x00000030
	MB_ICONASTERISK    = 0x00000040
	MB_ICONWARNING     = MB_ICONEXCLAMATION
	MB_ICONERROR       = MB_ICONHAND
*/
)

func myMessageBox(caption, text string, style uintptr) int {
	var nargs uintptr = 4
	ret, _, err := syscall.Syscall9(uintptr(messageBox),
		nargs,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		style,
		0,
		0,
		0,
		0,
		0)
	if err != 0 {
		log.Fatalf("Call MessageBox failed: %v", err)
	}

	return int(ret)
}

func main() {
	defer syscall.FreeLibrary(user32)
	ret := myMessageBox(
		"This is a Title",
		"This is a Message.",
		0x00000011) // MB_OKCANCEL|MB_ICONERROR
	log.Printf("Return: %d", ret)
}
