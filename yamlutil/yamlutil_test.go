package yamlutil

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/gofika/util/fileutil"
	"github.com/stretchr/testify/assert"
)

var (
	tempDir, _ = ioutil.TempDir("", "util")
)

type Foo struct {
	Name  string `yml:"name"`
	Value int    `yml:"value"`
}

func TestFile(t *testing.T) {
	foo := &Foo{
		Name:  "Jason",
		Value: 100,
	}
	tempName := path.Join(tempDir, "foo.yml")
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
	const indentData = `name: Jason
value: 100
`
	foo := &Foo{
		Name:  "Jason",
		Value: 100,
	}
	tempName := path.Join(tempDir, "foo.yml")
	err := WriteFileIndent(tempName, foo, 2)
	if !assert.Nil(t, err) {
		return
	}
	defer fileutil.Delete(tempName)
	data, err := ioutil.ReadFile(tempName)
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, string(data), indentData)
}
