package winapi

import (
	"syscall"
)

var (
	pPostMessageW = user32.NewProc("PostMessageW")
)

// PostMessageW places (posts) a message in the message queue.
func PostMessageW(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) error {
	ret, _, err := pPostMessageW.Call(
		uintptr(hwnd),
		uintptr(msg),
		wparam,
		lparam,
	)
	if ret == 0 {
		return err
	}
	return nil
}
