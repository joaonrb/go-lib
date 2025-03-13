package collection_test

import (
	"github.com/joaonrb/go-lib/collection"
	"testing"
)

func BenchmarkIntQueueCapacity10Add5(b *testing.B) {
	pushN(5, collection.NewQueue[int](10))
}

func BenchmarkIntQueueCapacity10Add10(b *testing.B) {
	pushN(10, collection.NewQueue[int](10))
}

func BenchmarkIntQueueCapacity100Add10(b *testing.B) {
	pushN(10, collection.NewQueue[int](100))
}

func BenchmarkIntQueueCapacity100Add25(b *testing.B) {
	pushN(25, collection.NewQueue[int](100))
}

func BenchmarkIntQueueCapacity100Add50(b *testing.B) {
	pushN(50, collection.NewQueue[int](100))
}

func BenchmarkIntQueueCapacity100Add100(b *testing.B) {
	pushN(100, collection.NewQueue[int](100))
}

func BenchmarkIntQueueCapacity1000Add100(b *testing.B) {
	pushN(500, collection.NewQueue[int](1000))
}

func BenchmarkIntQueueCapacity1000Add1000(b *testing.B) {
	pushN(1000, collection.NewQueue[int](1000))
}

func BenchmarkIntQueueCapacity10000Add1000(b *testing.B) {
	pushN(1000, collection.NewQueue[int](10000))
}

func BenchmarkIntQueueCapacity10000Add10000(b *testing.B) {
	pushN(10000, collection.NewQueue[int](10000))
}

func pushN(n int, queue *collection.Queue[int]) {
	for i := 0; i < n; i++ {
		queue.Push(i)
	}
}
