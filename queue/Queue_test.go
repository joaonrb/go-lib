package queue_test

import (
	"context"
	"fmt"
	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"

	"github.com/joaonrb/go-lib/queue"
)

func BenchmarkIntQueueCapacity10Add5(b *testing.B) {
	pushN(5, queue.New[int](10))
}

func BenchmarkIntQueueCapacity10Add10(b *testing.B) {
	pushN(10, queue.New[int](10))
}

func BenchmarkIntQueueCapacity100Add10(b *testing.B) {
	pushN(10, queue.New[int](100))
}

func BenchmarkIntQueueCapacity100Add25(b *testing.B) {
	pushN(25, queue.New[int](100))
}

func BenchmarkIntQueueCapacity100Add50(b *testing.B) {
	pushN(50, queue.New[int](100))
}

func BenchmarkIntQueueCapacity100Add100(b *testing.B) {
	pushN(100, queue.New[int](100))
}

func BenchmarkIntQueueCapacity1000Add100(b *testing.B) {
	pushN(500, queue.New[int](1000))
}

func BenchmarkIntQueueCapacity1000Add1000(b *testing.B) {
	pushN(1000, queue.New[int](1000))
}

func BenchmarkIntQueueCapacity10000Add1000(b *testing.B) {
	pushN(1000, queue.New[int](10000))
}

func BenchmarkIntQueueCapacity10000Add10000(b *testing.B) {
	pushN(10000, queue.New[int](10000))
}

func pushN(n int, queue *queue.Queue[int]) {
	for i := 0; i < n; i++ {
		queue.Push(i)
	}
	for !queue.IsEmpty() {
		queue.Pull()
	}
}

type QueueTestData struct {
	Name       string
	collection *queue.Queue[any]
}

var (
	data = []any{
		"carlos",
		10,
		10.11,
		true,
		"vanilla",
	}
	dataMap = map[any]bool{
		"carlos":  true,
		10:        true,
		10.11:     true,
		true:      true,
		"vanilla": true,
	}
)

func TestQueueShouldPullElementsWhenElementsWerePushedToTheCollection(t *testing.T) {
	tests := []QueueTestData{
		{"New collection size 5", queue.New[any](5)},
		{"New collection size 10", queue.New[any](10)},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			clt := test.collection
			for _, value := range data {
				clt.Push(value)
			}
			assert.False(t, clt.IsEmpty(), "IsEmpty is expected to be false")
			assert.Equalf(
				t,
				uint64(len(data)),
				clt.Length(),
				"Length is expected to be %d",
				len(data),
			)
			for !clt.IsEmpty() {
				result := testPull[any](t, clt)
				requireOK[any](t, result)
				result.WhenOK(func(value any) {
					require.Contains(
						t,
						dataMap,
						value,
						"Pull value %v is expected to be in data",
						value,
					)
				})
			}
		})
	}
}

func TestQueueMustPushShouldReturnFalseWhenCollectionIsFull(t *testing.T) {
	tests := []QueueTestData{
		{"New collection size 3", queue.New[any](3)},
	}
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			clt := test.collection
			clt.Push(1)
			clt.Push(2)
			clt.Push(3)
			result := clt.MustPush(4)
			requireOK[bool](t, result)
			result.WhenOK(func(value bool) {
				assert.False(t, value, "MustPush should have returned false")
			})
		})
	}
}

func TestQueueShouldLoopOverAllElementsInTheCollection(t *testing.T) {
	tests := []QueueTestData{
		{"New collection size 5", queue.New[any](5)},
		{"New collection size 10", queue.New[any](10)},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			clt := test.collection
			for _, value := range data {
				clt.Push(value)
			}
			clt.Close()
			loops := 0
			clt.ForEach(context.Background(), func(index int, value any) {
				assert.Contains(t, dataMap, value, "Value %v is expected to be in data", value)
				length := len(data) - index - 1
				assert.Equalf(t, length, int(clt.Length()), "Length is expected to be %d", length)
				loops++
			})
			assert.Equal(t, 0, int(clt.Length()), "Length is expected to be 0")
			assert.Equal(t, len(data), loops, "Loops is expected to be %d", len(data))
		})
	}
}

func TestQueueShouldLoopOverNElementsInTheCollection(t *testing.T) {
	tests := []QueueTestData{
		{"New collection size 5", queue.New[any](5)},
		{"New collection size 10", queue.New[any](10)},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			clt := test.collection
			for _, value := range data {
				clt.Push(value)
			}
			clt.Close()
			clt.ForN(3, func(index int, value any) {
				fmt.Println(index, value)
				assert.Contains(t, dataMap, value, "Value %v is expected to be in data", value)
				length := len(data) - index - 1
				assert.Equalf(t, length, int(clt.Length()), "Length is expected to be %d", length)
			})
			length := len(data) - 3
			assert.Equalf(t, length, int(clt.Length()), "Length is expected to be %d", length)
		})
	}
}

func testPull[T any](t *testing.T, queue *queue.Queue[T]) types.Result[T] {
	var result types.Result[T]
	c := make(chan types.Result[T])
	go func() {
		c <- queue.Pull()
	}()
	require.Eventually(t, func() bool {
		select {
		case result = <-c:
		default:
		}
		return result != nil
	}, time.Millisecond, 10*time.Microsecond, "Pull is expected to have a value")
	return result
}

func requireOK[T any](t *testing.T, result any) {
	require.IsTypef(t, types.OK[T]{}, result, "Pull result %s is expected to be OK[any]", result)
}
