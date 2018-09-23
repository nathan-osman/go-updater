package winapi

import (
	"unsafe"
)

var (
	pAdjustWindowRect = user32.NewProc("AdjustWindowRect")
)

// AdjustWindowRect calculates the required size of the window rectangle.
func AdjustWindowRect(rect *RECT, style uint32, menu bool) error {
	var menuInt int32
	if menu {
		menuInt = 1
	}
	ret, _, err := pAdjustWindowRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(style),
		uintptr(menuInt),
	)
	if ret == 0 {
		return err
	}
	return nil
}
