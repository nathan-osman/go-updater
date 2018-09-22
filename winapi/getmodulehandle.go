package winapi

import (
	"syscall"
)

var (
	pGetModuleHandleW = kernel32.NewProc("GetModuleHandleW")
)

// GetModuleHandle retrieves the application's module handle.
func GetModuleHandle() (syscall.Handle, error) {
	ret, _, err := pGetModuleHandleW.Call(uintptr(0))
	if ret == 0 {
		return 0, err
	}
	return syscall.Handle(ret), nil
}
