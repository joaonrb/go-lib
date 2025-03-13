package rand_test

import (
	go_rand "crypto/rand"
	"encoding/base64"
	"github.com/joaonrb/go-lib/rand"
	"math"
	"testing"
)

func defaultRandString(n int) string {
	buff := make([]byte, int(math.Ceil(float64(n)/1.33333333333)))
	_, _ = go_rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:n] // strip 1 extra character we get from odd length results
}

func BenchmarkDefaultString5(b *testing.B) {
	defaultRandString(5)
}

func BenchmarkString5(b *testing.B) {
	rand.String(5)
}

func BenchmarkSafeString5(b *testing.B) {
	rand.SafeString(5)
}

func BenchmarkDefaultString10(b *testing.B) {
	defaultRandString(10)
}

func BenchmarkString10(b *testing.B) {
	rand.String(10)
}

func BenchmarkSafeString10(b *testing.B) {
	rand.SafeString(10)
}

func BenchmarkDefaultString20(b *testing.B) {
	defaultRandString(20)
}

func BenchmarkString20(b *testing.B) {
	rand.String(20)
}

func BenchmarkSafeString20(b *testing.B) {
	rand.SafeString(20)
}

func BenchmarkDefaultString50(b *testing.B) {
	defaultRandString(50)
}

func BenchmarkString50(b *testing.B) {
	rand.String(50)
}

func BenchmarkSafeString50(b *testing.B) {
	rand.SafeString(50)
}

func BenchmarkDefaultString100(b *testing.B) {
	defaultRandString(100)
}

func BenchmarkString100(b *testing.B) {
	rand.String(100)
}

func BenchmarkSafeString100(b *testing.B) {
	rand.SafeString(100)
}

func BenchmarkDefaultString1000(b *testing.B) {
	defaultRandString(1000)
}

func BenchmarkString1000(b *testing.B) {
	rand.String(1000)
}

func BenchmarkSafeString1000(b *testing.B) {
	rand.SafeString(1000)
}

func BenchmarkDefaultString10000(b *testing.B) {
	defaultRandString(10000)
}

func BenchmarkString10000(b *testing.B) {
	rand.String(10000)
}

func BenchmarkSafeString10000(b *testing.B) {
	rand.SafeString(10000)
}
