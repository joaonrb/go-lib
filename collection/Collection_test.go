package collection_test

import (
	"context"
	"fmt"
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

func TestCollectionShouldPullElementsWhenElementsWerePushedToTheCollection(t *testing.T) {
	tests := []QueueTestData{
		{"NewQueue collection size 5", collection.NewQueue[any](5)},
		{"NewQueue collection size 10", collection.NewQueue[any](10)},
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

func TestCollectionMustPushShouldReturnFalseWhenCollectionIsFull(t *testing.T) {
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

func TestCollectionShouldLoopOverAllElementsInTheCollection(t *testing.T) {
	tests := []QueueTestData{
		{"NewQueue collection size 5", collection.NewQueue[any](5)},
		{"NewQueue collection size 10", collection.NewQueue[any](10)},
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

func TestCollectionShouldLoopOverNElementsInTheCollection(t *testing.T) {
	tests := []QueueTestData{
		{"NewQueue collection size 5", collection.NewQueue[any](5)},
		{"NewQueue collection size 10", collection.NewQueue[any](10)},
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
