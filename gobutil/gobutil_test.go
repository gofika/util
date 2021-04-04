package gobutil

import (
	. "gopkg.in/check.v1"
	"io/ioutil"
	"path"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type GobUtilSuite struct{}

var _ = Suite(&GobUtilSuite{})

var (
	tempDir, _ = ioutil.TempDir("", "util")
)

type Foo struct {
	Name  string
	Value int
}

func (s *GobUtilSuite) TestDeepCopy(c *C) {
	type Bar struct {
		Name  string
		Value int
	}
	foo := &Foo{"Jason", 100}
	var bar Bar
	err := DeepCopy(&bar, foo)
	c.Assert(err, IsNil)
	c.Assert(bar.Name, Equals, foo.Name)
	c.Assert(bar.Value, Equals, foo.Value)
}

func (s *GobUtilSuite) TestFile(c *C) {
	tempName := path.Join(tempDir, "foo.gob")
	foo := &Foo{"Jason", 100}

	err := WriteFile(tempName, foo)
	c.Assert(err, IsNil)
	var bar Foo
	err = ReadFile(tempName, &bar)
	c.Assert(err, IsNil)
	c.Assert(bar.Name, Equals, "Jason")
}
