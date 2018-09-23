package winapi

const (
	SM_CXSCREEN = 0
	SM_CYSCREEN = 1
)

var (
	pGetSystemMetrics = user32.NewProc("GetSystemMetrics")
)

// GetSystemMetrics retrieves the specified system metric.
func GetSystemMetrics(index int32) (int32, error) {
	ret, _, err := pGetSystemMetrics.Call(
		uintptr(index),
	)
	if ret == 0 {
		return 0, err
	}
	return int32(ret), nil
}
