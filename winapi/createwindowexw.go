package winapi

import (
	"syscall"
	"unsafe"
)

const (
	PROGRESS_CLASSW = "msctls_progress32"
)

const (
	SW_USE_DEFAULT = -1
)

const (
	WS_DLGFRAME = 0x00400000
	WS_CAPTION  = 0x00C00000
	WS_VISIBLE  = 0x10000000
	WS_CHILD    = 0x40000000
)

var (
	pCreateWindowExW = user32.NewProc("CreateWindowExW")
)

// CREATESTRUCTW defines the initialization parameters passed to the window procedure of an application.
type CREATESTRUCTW struct {
	CreateParams uintptr
	Instance     syscall.Handle
	Menu         syscall.Handle
	Parent       syscall.Handle
	Cy           int32
	Cx           int32
	Y            int32
	X            int32
	Style        uint32
	Name         uintptr
	Class        uintptr
	ExStyle      uint32
}

// CreateWindowExW creates an overlapped, pop-up, or child window.
func CreateWindowExW(
	className, windowName string,
	style uint32,
	x, y, width, height int32,
	parent, menu, instance syscall.Handle,
	param uintptr,
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
		param,
	)
	if ret == 0 {
		return 0, err
	}
	return syscall.Handle(ret), nil
}
