package winapi

import (
	"syscall"
	"unsafe"
)

var (
	pGetWindowRect = user32.NewProc("GetWindowRect")
)

// RECT defines the coordinates of the upper-left and lower-right corners of a rectangle.
type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

// GetWindowRect retrieves the dimensions of the bounding rectangle of the specified window.
func GetWindowRect(hwnd syscall.Handle, rect *RECT) error {
	ret, _, err := pGetWindowRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(rect)),
	)
	if ret == 0 {
		return err
	}
	return nil
}
