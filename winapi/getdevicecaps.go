package winapi

import (
	"syscall"
)

const (
	LOGPIXELSX = 88
	LOGPIXELSY = 90
)

var (
	pGetDeviceCaps = gdi32.NewProc("GetDeviceCaps")
)

// GetDeviceCaps retrieves device-specific information for the specified device.
func GetDeviceCaps(hdc syscall.Handle, index int32) int32 {
	ret, _, _ := pGetDeviceCaps.Call(
		uintptr(hdc),
		uintptr(index),
	)
	return int32(ret)
}
