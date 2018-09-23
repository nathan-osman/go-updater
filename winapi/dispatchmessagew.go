package winapi

import (
	"unsafe"
)

var (
	pDispatchMessageW = user32.NewProc("DispatchMessageW")
)

// DispatchMessageW dispatches a message to a window procedure.
func DispatchMessageW(msg *MSG) uintptr {
	ret, _, _ := pDispatchMessageW.Call(
		uintptr(unsafe.Pointer(msg)),
	)
	return ret
}
