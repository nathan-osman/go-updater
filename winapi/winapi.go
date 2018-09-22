package winapi

import (
	"syscall"
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")
)
