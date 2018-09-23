package winapi

import (
	"unsafe"
)

var (
	pTranslateMessage = user32.NewProc("TranslateMessage")
)

// TranslateMessage translates virtual-key messages into character messages.
func TranslateMessage(msg *MSG) bool {
	ret, _, _ := pTranslateMessage.Call(
		uintptr(unsafe.Pointer(msg)),
	)
	return int32(ret) != 0
}
