package winapi

import (
	"syscall"
	"unsafe"
)

const (
	// PROGRESS_CLASSW is the class name for progress bars.
	PROGRESS_CLASSW = "msctls_progress32"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	pRegisterClassExW = user32.NewProc("RegisterClassExW")
	pCreateWindowExW  = user32.NewProc("CreateWindowExW")
)

// WNDCLASSEXW contains window class information.
type WNDCLASSEXW struct {
	size       uint32
	style      uint32
	wndProc    uintptr
	clsExtra   int32
	wndExtra   int32
	instance   syscall.Handle
	icon       syscall.Handle
	cursor     syscall.Handle
	background syscall.Handle
	menuName   *uint16
	className  *uint16
	iconSm     syscall.Handle
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
