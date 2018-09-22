package winapi

import (
	"syscall"
	"unsafe"
)

var (
	pRegisterClassExW = user32.NewProc("RegisterClassExW")
)

// WNDCLASSEXW contains window class information.
type WNDCLASSEXW struct {
	size       uint32
	style      uint32
	wndProc    uintptr
	clsExtra   int32
	wndExtra   int32
	instance   syscall.Handle
	icon       syscall.Handle
	cursor     syscall.Handle
	background syscall.Handle
	menuName   *uint16
	className  *uint16
	iconSm     syscall.Handle
}

// RegisterClassW registers a window class for use in CreateWindowW.
func RegisterClassW(wcx *WNDCLASSEXW) (uint16, error) {
	ret, _, err := pRegisterClassExW.Call(
		uintptr(unsafe.Pointer(wcx)),
	)
	if ret == 0 {
		return 0, err
	}
	return uint16(ret), nil
}
