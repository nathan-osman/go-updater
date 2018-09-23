package winapi

import (
	"syscall"
)

const (
	SWP_NOSIZE = 0x0001
)

var (
	pSetWindowPos = user32.NewProc("SetWindowPos")
)

// SetWindowPos changes the size, position, and Z order of a window.
func SetWindowPos(
	hwnd, hwndInsertAfter syscall.Handle,
	x, y, cx, cy int32,
	flags uint32,
) error {
	ret, _, err := pSetWindowPos.Call(
		uintptr(hwnd),
		uintptr(hwndInsertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(flags),
	)
	if ret == 0 {
		return err
	}
	return nil
}
