package winapi

import (
	"syscall"
	"unsafe"
)

const (
	PROGRESS_CLASSW = "msctls_progress32"
)

var (
	pCreateWindowExW = user32.NewProc("CreateWindowExW")
)

// CreateWindowW creates an overlapped, pop-up, or child window.
func CreateWindowW(
	className, windowName string,
	style uint32,
	x, y, width, height int32,
	parent, menu, instance syscall.Handle,
) (syscall.Handle, error) {
	ret, _, err := pCreateWindowExW.Call(
		uintptr(0),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(className))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(windowName))),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(parent),
		uintptr(menu),
		uintptr(instance),
		uintptr(0),
	)
	if ret == 0 {
		return 0, err
	}
	return syscall.Handle(ret), nil
}
