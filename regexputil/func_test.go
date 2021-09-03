package regexputil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	bar, matched := Match(`Foo(.+)`, "Foobar")
	if !assert.True(t, matched) {
		return
	}
	assert.Equal(t, bar, "bar")
}

func TestIsMatch(t *testing.T) {
	assert.True(t, IsMatch(`Foo(.+)`, "Foobar"))
}
