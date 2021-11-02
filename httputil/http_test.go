package httputil

import (
	"io/ioutil"
	"net/http"
	"path"
	"testing"

	"github.com/gofika/util/fileutil"
	"github.com/gofika/util/jsonutil"
	"github.com/stretchr/testify/assert"
)

var (
	tempDir, _ = ioutil.TempDir("", "util")
)

func TestReadString(t *testing.T) {
	resp, err := http.Get("https://httpbin.org/get")
	if !assert.Nil(t, err) {
		return
	}

	body, err := ReadString(resp)
	if !assert.Nil(t, err) {
		return
	}
	assert.Contains(t, body, `"url": "https://httpbin.org/get"`)
	// c.Assert(strings.Contains(body, `"url": "https://httpbin.org/get"`), Equals, true)
}

func TestReadJson(t *testing.T) {
	resp, err := http.Get("https://httpbin.org/get")
	if !assert.Nil(t, err) {
		return
	}
	type ResponseBody struct {
		Url string `json:"url"`
	}
	var respBody ResponseBody
	err = ReadJson(resp, &respBody)
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, respBody.Url, "https://httpbin.org/get")
}

func TestReadFile(t *testing.T) {
	resp, err := http.Get("https://httpbin.org/get")
	if !assert.Nil(t, err) {
		return
	}
	type ResponseBody struct {
		Url string `json:"url"`
	}
	var respBody ResponseBody
	tempName := path.Join(tempDir, "foo.json")
	_, err = ReadFile(resp, tempName)
	if !assert.Nil(t, err) {
		return
	}
	defer fileutil.Delete(tempName)
	err = jsonutil.ReadFile(tempName, &respBody)
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, respBody.Url, "https://httpbin.org/get")
}
