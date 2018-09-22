package winapi

import (
	"syscall"
)

const (
	IDC_ARROW = 32512
)

var (
	pLoadCursorW = user32.NewProc("LoadCursorW")
)

// LoadCursorW loads the specified cursor.
func LoadCursorW(cursorName uint32) (syscall.Handle, error) {
	ret, _, err := pLoadCursorW.Call(
		uintptr(0),
		uintptr(uint16(cursorName)),
	)
	if ret == 0 {
		return 0, err
	}
	return syscall.Handle(ret), nil
}
