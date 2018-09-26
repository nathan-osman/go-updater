package winapi

var (
	pMulDiv = kernel32.NewProc("MulDiv")
)

// MulDiv multiplies two 32-bit values and then divides the 64-bit result by a third 32-bit value.
func MulDiv(number, numerator, denominator int32) int32 {
	ret, _, _ := pMulDiv.Call(
		uintptr(number),
		uintptr(numerator),
		uintptr(denominator),
	)
	return int32(ret)
}
