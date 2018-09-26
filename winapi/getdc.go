package winapi

import (
	"syscall"
)

var (
	pGetDC = user32.NewProc("GetDC")
)

// GetDC retrieves a handle to a device context for the client area of a specified window.
func GetDC(hwnd syscall.Handle) (syscall.Handle, error) {
	ret, _, err := pGetDC.Call(
		uintptr(hwnd),
	)
	if ret == 0 {
		return 0, err
	}
	return syscall.Handle(ret), nil
}
