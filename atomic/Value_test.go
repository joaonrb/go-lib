package atomic_test

import (
	"github.com/joaonrb/go-lib/atomic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueSetShouldSetValue(t *testing.T) {
	value := atomic.NewValue(55)
	assert.Equal(t, 55, value.Value())
}

func TestValueSwapShouldSetANewValueAndReturnTheOldOne(t *testing.T) {
	value := atomic.NewValue(55)
	assert.Equal(t, 55, value.Swap(23))
	assert.Equal(t, 23, value.Value())
}

func TestValueSetShouldSetValueWithoutRaceCondition(t *testing.T) {
	value := atomic.NewValue(55)
	for i := 0; i < 100; i++ {
		go func() {
			value.Set(i)
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			assert.IsType(t, i, value.Value())
		}()
	}
}
