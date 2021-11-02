package gobutil

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
	Name  string
	Value int
}

func TestDeepCopy(t *testing.T) {
	type Bar struct {
		Name  string
		Value int
	}
	foo := &Foo{"Jason", 100}
	var bar Bar
	err := DeepCopy(&bar, foo)
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, bar.Name, foo.Name)
	assert.Equal(t, bar.Value, foo.Value)
}

func TestFile(t *testing.T) {
	tempName := path.Join(tempDir, "foo.gob")
	foo := &Foo{"Jason", 100}

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
	assert.Equal(t, bar.Name, "Jason")
}
