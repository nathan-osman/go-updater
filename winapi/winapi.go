package winapi

import (
	"syscall"
)

var (
	gdi32    = syscall.NewLazyDLL("gdi32.dll")
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	user32   = syscall.NewLazyDLL("user32.dll")
)

// LOWORD retrieves the low-order word from the specified value.
func LOWORD(v uint32) uint16 {
	return uint16(v)
}
