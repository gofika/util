package fileutil

import (
	. "gopkg.in/check.v1"
	"io/ioutil"
	"path"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type FileUtilSuite struct{}

var _ = Suite(&FileUtilSuite{})

var (
	tempDir, _ = ioutil.TempDir("", "util")
)

func (s *FileUtilSuite) TestWriteFile(c *C) {
	tempName := path.Join(tempDir, "foo.txt")
	data := []byte("Hello")
	err := WriteFile(tempName, data)
	c.Assert(err, IsNil)
	defer Delete(tempName)
	c.Assert(IsExist(tempName), Equals, true)
	rData, err := ioutil.ReadFile(tempName)
	c.Assert(err, IsNil)
	c.Assert(data, DeepEquals, rData)
}

func (s *FileUtilSuite) TestDeleteAll(c *C) {
	tempName := path.Join(tempDir, "/foo/bar/baz.js")
	f, err := OpenWrite(tempName)
	c.Assert(err, IsNil)
	defer DeleteAll(tempDir)
	err = f.Close()
	c.Assert(err, IsNil)
	DeleteAll(path.Join(tempDir, "/foo"))
	c.Assert(IsExist(path.Join(tempDir, "/foo/bar")), Equals, false)
}

func (s *FileUtilSuite) TestCurrentDir(c *C) {
	c.Assert(CurrentDir(), NotNil)
}
