package winapi

import (
	"syscall"
	"unsafe"
)

const (
	COLOR_WINDOW = 5
)

var (
	pRegisterClassExW = user32.NewProc("RegisterClassExW")
)

// WNDCLASSEXW contains window class information.
type WNDCLASSEXW struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   syscall.Handle
	Icon       syscall.Handle
	Cursor     syscall.Handle
	Background syscall.Handle
	MenuName   *uint16
	ClassName  *uint16
	IconSm     syscall.Handle
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
