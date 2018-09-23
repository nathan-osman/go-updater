package winapi

import (
	"syscall"
)

const (
	SW_SHOW = 5
)

var (
	pShowWindow = user32.NewProc("ShowWindow")
)

// ShowWindow sets the specified window's show state.
func ShowWindow(hwnd syscall.Handle, cmdShow int32) bool {
	ret, _, _ := pShowWindow.Call(
		uintptr(hwnd),
		uintptr(cmdShow),
	)
	return ret != 0
}
