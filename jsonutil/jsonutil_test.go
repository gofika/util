package jsonutil

import (
	"github.com/leaker/util/fileutil"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"path"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type JsonUtilSuite struct{}

var _ = Suite(&JsonUtilSuite{})

var (
	tempDir, _ = ioutil.TempDir("", "util")
)

type Foo struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (s *JsonUtilSuite) TestFile(c *C) {
	foo := &Foo{
		Name:  "Jason",
		Value: 100,
	}
	tempName := path.Join(tempDir, "foo.json")
	err := WriteFile(tempName, foo)
	c.Assert(err, IsNil)
	defer fileutil.Delete(tempName)
	var bar Foo
	err = ReadFile(tempName, &bar)
	c.Assert(err, IsNil)
	c.Assert(&bar, DeepEquals, foo)
}

func (s *JsonUtilSuite) TestSaveFileIndent(c *C) {
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
	c.Assert(err, IsNil)
	defer fileutil.Delete(tempName)
	data, err := ioutil.ReadFile(tempName)
	c.Assert(err, IsNil)
	c.Assert(string(data), Equals, indentData)
}
