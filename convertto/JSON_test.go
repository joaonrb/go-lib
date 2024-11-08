package convertto_test

import (
	"testing"
	"time"

	"github.com/joaonrb/go-lib/convertto"
	"github.com/joaonrb/go-lib/types"
	"github.com/stretchr/testify/assert"
)

func TestJSONShouldReturnAJsonLikeStringWhenValueIsAString(t *testing.T) {
	assertOK(t, `"João Nuno"`, convertto.JSON("João Nuno"))
}

func TestJSONShouldReturnAJsonLikeNumberWhenValueIsANumber(t *testing.T) {
	assertOK(t, `10`, convertto.JSON(10))
}

func TestJSONShouldReturnADatetimeWhenValueIsADatetime(t *testing.T) {
	assertOK(
		t,
		`"2024-10-19T12:02:00Z"`,
		convertto.JSON(time.Date(2024, 10, 19, 12, 2, 0, 0, time.UTC)),
	)
}

type Person struct {
	Name     string
	Age      int
	Birthday time.Time
}

func TestJSONShouldReturnAObjectWhenValueIsAnObject(t *testing.T) {
	person := Person{
		Name:     "David Doe",
		Age:      28,
		Birthday: time.Date(1992, 10, 19, 12, 2, 0, 0, time.UTC),
	}
	assertOK(
		t,
		`{"Name":"David Doe","Age":28,"Birthday":"1992-10-19T12:02:00Z"}`,
		convertto.JSON(person),
	)
}

func TestPrettyJSONShouldReturnAObjectWhenValueIsAnObject(t *testing.T) {
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
	assertOK(t, expected, convertto.PrettyJSON(person))
}

func assertOK[T any](t *testing.T, expected any, result types.Result[T]) {
	assert.IsType(t, types.OK[T]{}, result, "result not from type OK")
	assert.Equal(t, expected, result.(types.OK[T]).Value, "value does not match the expected")
}
