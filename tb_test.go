package tb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoalesce(t *testing.T) {
	assert.Equal(t, "hello", Coalesce("", "hello"))
	assert.Equal(t, "", Coalesce("", ""))
	a := "hello"
	assert.Equal(t, &a, Coalesce(nil, &a))
}

func TestZeroIfNil(t *testing.T) {
	a := "hello"
	assert.Equal(t, "hello", ZeroIfNil(&a))
	assert.Equal(t, "", ZeroIfNil[string](nil))
}

func TestIf(t *testing.T) {
	assert.Equal(t, "hello", If(true, "hello", "world"))
	assert.Equal(t, "world", If(false, "hello", "world"))
}
