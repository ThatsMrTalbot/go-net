package http2

import "runtime"

// Wipe takes a buffer and wipes it with zeroes.
func wipeBuffer(buf []byte) {
	for i := range buf {
		buf[i] = 0
	}

	// This should keep buf's backing array live and thus prevent dead store
	// elimination, according to discussion at
	// https://github.com/golang/go/issues/33325 .
	runtime.KeepAlive(buf)
}
