package gobutil

import (
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type GobUtilSuite struct{}

var _ = Suite(&GobUtilSuite{})

func (s *GobUtilSuite) TestDeepCopy(c *C) {
	type Foo struct {
		Name  string
		Value int
	}
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
