package winapi

var (
	pPostQuitMessage = user32.NewProc("PostQuitMessage")
)

// PostQuitMessage indicates to the system that a thread has made a request to terminate.
func PostQuitMessage(exitCode int32) {
	pPostQuitMessage.Call(uintptr(exitCode))
}
