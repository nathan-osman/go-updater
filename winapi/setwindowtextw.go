package winapi

import (
	"syscall"
	"unsafe"
)

var (
	pSetWindowTextW = user32.NewProc("SetWindowTextW")
)

// SetWindowTextW changes the text of the specified window's title bar.
func SetWindowTextW(hwnd syscall.Handle, text string) error {
	ret, _, err := pSetWindowTextW.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
	)
	if ret == 0 {
		return err
	}
	return nil
}
