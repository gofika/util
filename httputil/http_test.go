package httputil

import (
	"github.com/leaker/util/fileutil"
	"github.com/leaker/util/jsonutil"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type HttpUtilSuite struct{}

var _ = Suite(&HttpUtilSuite{})

var (
	tempDir, _ = ioutil.TempDir("", "util")
)

func (s *HttpUtilSuite) TestReadString(c *C) {
	resp, err := http.Get("https://httpbin.org/get")
	c.Assert(err, IsNil)

	body, err := ReadString(resp)
	c.Assert(err, IsNil)
	c.Assert(strings.Contains(body, `"url": "https://httpbin.org/get"`), Equals, true)
}

func (s *HttpUtilSuite) TestReadJson(c *C) {
	resp, err := http.Get("https://httpbin.org/get")
	c.Assert(err, IsNil)
	type ResponseBody struct {
		Url string `json:"url"`
	}
	var respBody ResponseBody
	err = ReadJson(resp, &respBody)
	c.Assert(err, IsNil)
	c.Assert(respBody.Url, Equals, "https://httpbin.org/get")
}

func (s *HttpUtilSuite) TestReadFile(c *C) {
	resp, err := http.Get("https://httpbin.org/get")
	c.Assert(err, IsNil)
	type ResponseBody struct {
		Url string `json:"url"`
	}
	var respBody ResponseBody
	tempName := path.Join(tempDir, "foo.json")
	_, err = ReadFile(resp, tempName)
	c.Assert(err, IsNil)
	defer fileutil.Delete(tempName)
	err = jsonutil.ReadFile(tempName, &respBody)
	c.Assert(err, IsNil)
	c.Assert(respBody.Url, Equals, "https://httpbin.org/get")
}
