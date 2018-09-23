package winapi

import (
	"syscall"
)

const (
	PBM_SETPOS = WM_USER + 2
)

var (
	pSendMessageW = user32.NewProc("SendMessageW")
)

// SendMessageW sends the specified message to a window or windows.
func SendMessageW(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	ret, _, _ := pSendMessageW.Call(
		uintptr(hwnd),
		uintptr(msg),
		wparam,
		lparam,
	)
	return ret
}
