package fileutil

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	tempDir, _ = os.MkdirTemp("", "util")
)

func TestWriteFile(t *testing.T) {
	tempName := path.Join(tempDir, "foo.txt")
	data := []byte("Hello")
	err := WriteFile(tempName, data)
	if !assert.Nil(t, err) {
		return
	}
	defer Delete(tempName)
	if !assert.True(t, IsExist(tempName)) {
		return
	}
	rData, err := os.ReadFile(tempName)
	if !assert.Nil(t, err) {
		return
	}
	assert.EqualValues(t, data, rData)
}

func TestDeleteAll(t *testing.T) {
	tempName := path.Join(tempDir, "/foo/bar/baz.js")
	f, err := OpenWrite(tempName)
	if !assert.Nil(t, err) {
		return
	}
	defer DeleteAll(tempDir)
	err = f.Close()
	if !assert.Nil(t, err) {
		return
	}
	DeleteAll(path.Join(tempDir, "/foo"))
	assert.False(t, IsExist(path.Join(tempDir, "/foo/bar")))
}

func TestCurrentDir(t *testing.T) {
	assert.True(t, CurrentDir() != "")
}
