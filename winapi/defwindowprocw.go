package winapi

import (
	"syscall"
)

const (
	WM_CREATE  = 0x0001
	WM_DESTROY = 0x0002
	WM_CLOSE   = 0x0010
)

var (
	pDefWindowProcW = user32.NewProc("DefWindowProcW")
)

// DefWindowProcW calls the default window procedure.
func DefWindowProcW(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	ret, _, _ := pDefWindowProcW.Call(
		uintptr(hwnd),
		uintptr(msg),
		uintptr(wparam),
		uintptr(lparam),
	)
	return uintptr(ret)
}
