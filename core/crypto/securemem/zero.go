package securemem

// Zeroize zeros out a byte slice.
func Zeroize(b []byte) {
	// Note: Go's garbage collector makes true zeroization challenging.
	// This is a best-effort implementation.
	for i := range b {
		b[i] = 0
	}
}

// ZeroizeAfterFn executes a function and zeros out the returned byte slice.
func ZeroizeAfterFn(fn func() []byte) []byte {
	result := fn()
	Zeroize(result)
	return result
}