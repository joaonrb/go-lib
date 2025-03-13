package rand

import (
	"math/rand"
	"time"
	"unsafe"
)

const (
	charset     = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	safecharset = "3456789fghjmpqvwxyzFGHJMPQVWXYZ"
)
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	safeLetterIdxBits = 5                        // 6 bits to represent a letter index
	safeLetterIdxMask = 1<<safeLetterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	safeLetterIdxMax  = 31 / safeLetterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// String Copied from https://stackoverflow.com/a/31832326/2521743 and is credited to the original author.
func String(n int) string {

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(charset) {
			b[i] = charset[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

func SafeString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), safeLetterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), safeLetterIdxMax
		}
		if idx := int(cache & safeLetterIdxMask); idx < len(safecharset) {
			b[i] = safecharset[idx]
			i--
		}
		cache >>= safeLetterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
