package convert_test

import (
	"testing"
	"time"

	"github.com/joaonrb/go-lib/convert"

	"github.com/joaonrb/go-lib/monad"
	"github.com/stretchr/testify/assert"
)

func TestToJSONShouldReturnAJsonLikeStringWhenValueIsAString(t *testing.T) {
	assertOK(t, `"João Nuno"`, convert.ToJSON("João Nuno"))
}

func TestToJSONShouldReturnAJsonLikeNumberWhenValueIsANumber(t *testing.T) {
	assertOK(t, `10`, convert.ToJSON(10))
}

func TestToJSONShouldReturnADatetimeWhenValueIsADatetime(t *testing.T) {
	assertOK(
		t,
		`"2024-10-19T12:02:00Z"`,
		convert.ToJSON(time.Date(2024, 10, 19, 12, 2, 0, 0, time.UTC)),
	)
}

type Person struct {
	Name     string
	Age      int
	Birthday time.Time
}

func TestToJSONShouldReturnAObjectWhenValueIsAnObject(t *testing.T) {
	person := Person{
		Name:     "David Doe",
		Age:      28,
		Birthday: time.Date(1992, 10, 19, 12, 2, 0, 0, time.UTC),
	}
	assertOK(
		t,
		`{"Name":"David Doe","Age":28,"Birthday":"1992-10-19T12:02:00Z"}`,
		convert.ToJSON(person),
	)
}

func TestToPrettyJSONShouldReturnAObjectWhenValueIsAnObject(t *testing.T) {
	person := Person{
		Name:     "David Doe",
		Age:      28,
		Birthday: time.Date(1992, 10, 19, 12, 2, 0, 0, time.UTC),
	}
	expected := `{
    "Name": "David Doe",
    "Age": 28,
    "Birthday": "1992-10-19T12:02:00Z"
}`
	assertOK(t, expected, convert.ToPrettyJSON(person))
}

func assertOK[T any](t *testing.T, expected any, result monad.Result[T]) {
	assert.IsType(t, monad.OK[T]{}, result, "result not from type OK")
	assert.Equal(t, expected, result.(monad.OK[T]).Value, "value does not match the expected")
}
