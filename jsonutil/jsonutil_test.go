package jsonutil

import (
	"os"
	"path"
	"testing"

	"github.com/gofika/util/fileutil"
	"github.com/stretchr/testify/assert"
)

var (
	tempDir, _ = os.MkdirTemp("", "util")
)

type Foo struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func TestFile(t *testing.T) {
	foo := &Foo{
		Name:  "Jason",
		Value: 100,
	}
	tempName := path.Join(tempDir, "foo.json")
	err := WriteFile(tempName, foo)
	if !assert.Nil(t, err) {
		return
	}
	defer fileutil.Delete(tempName)
	var bar Foo
	err = ReadFile(tempName, &bar)
	if !assert.Nil(t, err) {
		return
	}
	assert.EqualValues(t, &bar, foo)
}

func TestSaveFileIndent(t *testing.T) {
	const indentData = `{
    "name": "Jason",
    "value": 100
}
`
	foo := &Foo{
		Name:  "Jason",
		Value: 100,
	}
	tempName := path.Join(tempDir, "foo.json")
	err := WriteFileIndent(tempName, foo, "    ")
	if !assert.Nil(t, err) {
		return
	}
	defer fileutil.Delete(tempName)
	data, err := os.ReadFile(tempName)
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, string(data), indentData)
}
