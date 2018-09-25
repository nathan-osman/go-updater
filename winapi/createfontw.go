package winapi

import (
	"syscall"
	"unsafe"
)

var (
	pCreateFontW = user32.NewProc("CreateFontW")
)

// CreateFontW creates a logical font with the specified characteristics.
func CreateFontW(
	height, width, escapement, orientation, weight int32,
	italic, underline, strikeOut,
	charSet, outPrecision, clipPrecision, quality, pitchAndFamily uint32,
	faceName string,
) (syscall.Handle, error) {
	ret, _, err := pCreateFontW.Call(
		uintptr(height),
		uintptr(width),
		uintptr(escapement),
		uintptr(orientation),
		uintptr(weight),
		uintptr(italic),
		uintptr(underline),
		uintptr(strikeOut),
		uintptr(charSet),
		uintptr(outPrecision),
		uintptr(clipPrecision),
		uintptr(quality),
		uintptr(pitchAndFamily),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(faceName))),
	)
	if ret == 0 {
		return 0, err
	}
	return syscall.Handle(ret), nil
}
