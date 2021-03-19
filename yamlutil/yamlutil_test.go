package yamlutil

import (
	"github.com/leaker/util/fileutil"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"path"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type YamlUtilSuite struct{}

var _ = Suite(&YamlUtilSuite{})

var (
	tempDir, _ = ioutil.TempDir("", "util")
)

type Foo struct {
	Name  string `yml:"name"`
	Value int    `yml:"value"`
}

func (s *YamlUtilSuite) TestFile(c *C) {
	foo := &Foo{
		Name:  "Jason",
		Value: 100,
	}
	tempName := path.Join(tempDir, "foo.yml")
	err := WriteFile(tempName, foo)
	c.Assert(err, IsNil)
	defer fileutil.Delete(tempName)
	var bar Foo
	err = ReadFile(tempName, &bar)
	c.Assert(err, IsNil)
	c.Assert(&bar, DeepEquals, foo)
}

func (s *YamlUtilSuite) TestSaveFileIndent(c *C) {
	const indentData = `name: Jason
value: 100
`
	foo := &Foo{
		Name:  "Jason",
		Value: 100,
	}
	tempName := path.Join(tempDir, "foo.yml")
	err := WriteFileIndent(tempName, foo, 2)
	c.Assert(err, IsNil)
	defer fileutil.Delete(tempName)
	data, err := ioutil.ReadFile(tempName)
	c.Assert(err, IsNil)
	c.Assert(string(data), Equals, indentData)
}
