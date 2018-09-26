package winapi

import (
	"syscall"
)

var (
	pReleaseDC = user32.NewProc("ReleaseDC")
)

// ReleaseDC releases a device context.
func ReleaseDC(hwnd, hdc syscall.Handle) bool {
	ret, _, _ := pReleaseDC.Call(
		uintptr(hwnd),
		uintptr(hdc),
	)
	return ret != 0
}
