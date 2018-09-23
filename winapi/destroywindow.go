package winapi

import "syscall"

var (
	pDestroyWindow = user32.NewProc("DestroyWindow")
)

// DestroyWindow destroys the specified window.
func DestroyWindow(hwnd syscall.Handle) error {
	ret, _, err := pDestroyWindow.Call(uintptr(hwnd))
	if ret == 0 {
		return err
	}
	return nil
}
