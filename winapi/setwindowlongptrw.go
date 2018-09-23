package winapi

import (
	"syscall"
)

var (
	pSetWindowLongPtrW = user32.NewProc("SetWindowLongPtrW")
)

// SetWindowLongPtrW changes an attribute of the specified window.
func SetWindowLongPtrW(hwnd syscall.Handle, index int32, newLong uintptr) uintptr {
	ret, _, _ := pSetWindowLongPtrW.Call(
		uintptr(hwnd),
		uintptr(index),
		newLong,
	)
	return uintptr(ret)
}
