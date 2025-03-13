package collection_test

import (
	"testing"
	"time"

	"github.com/joaonrb/go-lib/collection"
	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type QueueTestData struct {
	Name       string
	collection collection.Collection[any]
}

var data = map[any]bool{
	"carlos":  true,
	10:        true,
	10.11:     true,
	true:      true,
	"vanilla": true,
}

func TestCollectionShouldReadElementsWhenElementsArePutOnTheCollection(t *testing.T) {
	tests := []QueueTestData{
		{"NewQueue collection size 5", collection.NewQueue[any](5)},
		{"NewQueue collection size 10", collection.NewQueue[any](10)},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			clt := test.collection
			for value := range data {
				clt.Push(value)
			}
			assert.False(t, clt.IsEmpty(), "IsEmpty is expected to be false")
			assert.Equalf(t, len(data), clt.Length(), "Length is expected to be %d", len(data))
			for !clt.IsEmpty() {
				result := testPull[any](t, clt)
				requireOK[any](t, result)
				result.WhenOK(func(value any) {
					require.Contains(
						t,
						data,
						value,
						"Pull value %v is expected to be in data",
						value,
					)
				})
			}
		})
	}
}

func TestCollectionMustPutShouldReturnFalseWhenCollectionIsFull(t *testing.T) {
	tests := []QueueTestData{
		{"NewQueue collection size 3", collection.NewQueue[any](3)},
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

func testPull[T any](t *testing.T, clt collection.Collection[T]) types.Result[T] {
	var result types.Result[T]
	c := make(chan types.Result[T])
	go func() {
		c <- clt.Pull()
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
