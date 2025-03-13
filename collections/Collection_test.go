package collections_test

import (
	"github.com/joaonrb/go-lib/collections"
	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type CollectionTestData struct {
	Name       string
	Collection collections.Collection[any]
}

var data = map[any]bool{
	"carlos":  true,
	10:        true,
	10.11:     true,
	true:      true,
	"vanilla": true,
}

func TestCollectionShouldReadElementsWhenElementsArePutOnTheCollection(t *testing.T) {
	tests := []CollectionTestData{
		{"Queue Collection size 4", collections.Queue[any](4)},
		{"Queue Collection size 10", collections.Queue[any](10)},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			collection := test.Collection
			for value := range data {
				collection.Push(value)
			}
			assert.False(t, collection.IsEmpty(), "IsEmpty is expected to be false")
			assert.Equalf(t, len(data), collection.Length(), "Length is expected to be %d", len(data))
			for !collection.IsEmpty() {
				result := testPop(t, collection)
				requireOK[any](t, result)
				result.WhenOK(func(value any) {
					assert.Contains(t, data, value, "Pop value %v is expected to be in data", value)
				})
			}
		})
	}
}

func TestCollectionMustPutShouldReturnFalseWhenCollectionIsFull(t *testing.T) {
	tests := []CollectionTestData{
		{"Queue Collection size 3", collections.Queue[any](3)},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			collection := test.Collection
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

func testPop[T any](t *testing.T, queue collections.Collection[T]) types.Result[T] {
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
