package winapi

import (
	"syscall"
	"unsafe"
)

var (
	pGetMessageW = user32.NewProc("GetMessageW")
)

// POINT defines the x and y coordinates of a point.
type POINT struct {
	X, Y int32
}

// MSG contains message information from a thread's message queue.
type MSG struct {
	Hwnd    syscall.Handle
	Message uint32
	Wparam  uintptr
	Lparam  uintptr
	Time    uint32
	Pt      POINT
}

// GetMessageW retrieves a message from the calling thread's message queue.
func GetMessageW(msg *MSG, hwnd syscall.Handle, msgFilterMin, msgFilterMax uint32) (bool, error) {
	ret, _, err := pGetMessageW.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
	)
	if int32(ret) == -1 {
		return false, err
	}
	return int32(ret) != 0, nil
}
