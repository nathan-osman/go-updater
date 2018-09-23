// +build windows

package dialog

import (
	"syscall"
	"unsafe"

	"github.com/nathan-osman/go-updater/winapi"
)

const (
	className = "updateDialog"

	ID_BUTTON = 101
)

// The call to CreateWindowW includes a pointer to the WindowsDialog, which is
// sent as a parameter during WM_CREATE; this can be used by subsequent calls
// to the window procedure to retrieve the pointer

func wndProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	dlg := (*WindowsDialog)(unsafe.Pointer(
		winapi.GetWindowLongPtrW(hwnd, winapi.GWLP_USERDATA),
	))
	if dlg == nil && msg == winapi.WM_CREATE {
		dlg = (*WindowsDialog)(unsafe.Pointer(lparam))
		winapi.SetWindowLongPtrW(hwnd, winapi.GWLP_USERDATA, lparam)
	}
	if dlg != nil {
		return dlg.wndProc(hwnd, msg, wparam, lparam)
	}
	return winapi.DefWindowProcW(hwnd, msg, wparam, lparam)
}

func init() {
	var (
		instance, _ = winapi.GetModuleHandle()
		cursor, _   = winapi.LoadCursorW(winapi.IDC_ARROW)
		wcx         = winapi.WNDCLASSEXW{
			Size:       uint32(unsafe.Sizeof(winapi.WNDCLASSEXW{})),
			WndProc:    syscall.NewCallback(wndProc),
			Instance:   instance,
			Cursor:     cursor,
			Background: winapi.COLOR_WINDOW,
			ClassName:  syscall.StringToUTF16Ptr(className),
		}
	)
	winapi.RegisterClassW(&wcx)
}

// WindowsDialog implements the Dialog interface using the Windows API.
type WindowsDialog struct {
	hwnd     syscall.Handle
	cancelCh chan<- bool
}

func (w *WindowsDialog) wndProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case winapi.WM_COMMAND:
		if winapi.LOWORD(uint32(wparam)) == ID_BUTTON {
			close(w.cancelCh)
		}
	case winapi.WM_DESTROY:
		winapi.PostQuitMessage(0)
	default:
		return winapi.DefWindowProcW(hwnd, msg, wparam, lparam)
	}
	return 0
}

func (w *WindowsDialog) initialize() {
	w.hwnd, _ = winapi.CreateWindowExW(
		className,
		"Software Update",
		winapi.WS_DLGFRAME|winapi.WS_CAPTION|winapi.WS_VISIBLE,
		winapi.SW_USE_DEFAULT,
		winapi.SW_USE_DEFAULT,
		winapi.SW_USE_DEFAULT,
		winapi.SW_USE_DEFAULT,
		0, 0, 0,
		uintptr(unsafe.Pointer(w)),
	)
	winapi.CreateWindowExW(
		"STATIC",
		"Initializing...",
		winapi.WS_CHILD|winapi.WS_VISIBLE,
		10, 10, 380, 100,
		w.hwnd,
		0, 0, 0,
	)
	winapi.CreateWindowExW(
		"BUTTON",
		"Cancel",
		winapi.WS_CHILD|winapi.WS_VISIBLE,
		310, 80, 80, 30,
		w.hwnd,
		ID_BUTTON,
		0, 0,
	)
	winapi.CreateWindowExW(
		winapi.PROGRESS_CLASSW,
		"",
		winapi.WS_CHILD|winapi.WS_VISIBLE,
		10, 40, 380, 30,
		w.hwnd,
		0, 0, 0,
	)
}

func (w *WindowsDialog) resizeAndCenter() {
	var (
		screenWidth, _  = winapi.GetSystemMetrics(winapi.SM_CXSCREEN)
		screenHeight, _ = winapi.GetSystemMetrics(winapi.SM_CYSCREEN)
		rect            = winapi.RECT{
			Right:  400,
			Bottom: 120,
		}
	)
	winapi.AdjustWindowRect(&rect, winapi.WS_DLGFRAME|winapi.WS_CAPTION, false)
	var (
		dialogWidth  = rect.Right - rect.Left
		dialogHeight = rect.Bottom - rect.Top
	)
	winapi.SetWindowPos(
		w.hwnd,
		winapi.HWND_TOP,
		screenWidth/2-dialogWidth/2,
		screenHeight/2-dialogHeight/2,
		dialogWidth, dialogHeight,
		0,
	)
}

// New creates a new Windows dialog.
func New() Dialog {
	return &WindowsDialog{}
}

// Exec shows the window and runs an event loop.
func (w *WindowsDialog) Exec(cancelCh chan<- bool) {
	w.cancelCh = cancelCh
	w.initialize()
	w.resizeAndCenter()
	for {
		var msg winapi.MSG
		if m, _ := winapi.GetMessageW(&msg, 0, 0, 0); m {
			winapi.TranslateMessage(&msg)
			winapi.DispatchMessageW(&msg)
		} else {
			break
		}
	}
}

// SetStatus sets the text of the status label.
func (w *WindowsDialog) SetStatus(string) {
	//...
}

// SetProgress sets the value of the progress bar.
func (w *WindowsDialog) SetProgress(int) {
	//...
}

// Close destroys the dialog and causes .
func (w *WindowsDialog) Close() {
	winapi.DestroyWindow(w.hwnd)
}
