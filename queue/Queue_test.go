package queue_test

import (
	"testing"
	"time"

	"github.com/joaonrb/go-lib/queue"
	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type QueueTestData struct {
	Name  string
	Queue *queue.Queue[any]
}

var data = map[any]bool{
	"carlos":  true,
	10:        true,
	10.11:     true,
	true:      true,
	"vanilla": true,
}

func TestQueueShouldReadElementsWhenElementsArePutOnTheQueue(t *testing.T) {
	tests := []QueueTestData{
		{"New Collection size 5", queue.New[any](5)},
		{"New Collection size 10", queue.New[any](10)},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			queue := test.Queue
			for value := range data {
				queue.Push(value)
			}
			assert.False(t, queue.IsEmpty(), "IsEmpty is expected to be false")
			assert.Equalf(t, len(data), queue.Length(), "Length is expected to be %d", len(data))
			for !queue.IsEmpty() {
				result := testPop[any](t, queue)
				requireOK[any](t, result)
				result.WhenOK(func(value any) {
					assert.Contains(t, data, value, "Pop value %v is expected to be in data", value)
				})
			}
		})
	}
}

func TestCollectionMustPutShouldReturnFalseWhenCollectionIsFull(t *testing.T) {
	tests := []QueueTestData{
		{"New Collection size 3", queue.New[any](3)},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			collection := test.Queue
			collection.Push(1)
			collection.Push(2)
			collection.Push(3)
			result := collection.MustPush(4)
			requireOK[bool](t, result)
			result.WhenOK(func(value bool) {
				assert.False(t, value, "MustPush should have returned false")
			})
		})
	}
}

func testPop[T any](t *testing.T, queue *queue.Queue[T]) types.Result[T] {
	var result types.Result[T]
	c := make(chan types.Result[T])
	go func() {
		c <- queue.Pop()
	}()
	require.Eventually(t, func() bool {
		select {
		case result = <-c:
		default:
		}
		return result != nil
	}, time.Millisecond, 10*time.Microsecond, "Pop is expected to have a value")
	return result
}

func requireOK[T any](t *testing.T, result any) {
	require.IsTypef(t, types.OK[T]{}, result, "Pop result %s is expected to be OK[any]", result)
}
