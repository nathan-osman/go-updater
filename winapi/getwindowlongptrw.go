package winapi

import (
	"syscall"
)

const (
	GWLP_USERDATA = -21
)

var (
	pGetWindowLongPtrW = user32.NewProc("GetWindowLongPtrW")
)

// GetWindowLongPtrW retrieves information about the specified window.
func GetWindowLongPtrW(hwnd syscall.Handle, index int32) uintptr {
	ret, _, _ := pGetWindowLongPtrW.Call(
		uintptr(hwnd),
		uintptr(index),
	)
	return ret
}
