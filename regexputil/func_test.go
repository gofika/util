package regexputil

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type RegexpUtilSuite struct{}

var _ = Suite(&RegexpUtilSuite{})

func (s *RegexpUtilSuite) TestMatch(c *C) {
	c.Assert(Match(`Foo(.+)`, "Foobar"), Equals, "bar")
}

func (s *RegexpUtilSuite) TestIsMatch(c *C) {
	c.Assert(IsMatch(`Foo(.+)`, "Foobar"), Equals, true)
}
