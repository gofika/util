package httputil

import (
	"encoding/json"
	"github.com/leaker/util/fileutil"
	"github.com/leaker/util/regexputil"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// ReadAll read all data from resp and close resp.Body
func ReadAll(resp *http.Response) (b []byte, err error) {
	defer func() {
		err = resp.Body.Close()
	}()
	var bodyReader io.Reader
	bodyReader = resp.Body
	// auto convert encoding
	contentType := resp.Header.Get("Content-Type")
	contentCharset, matched := regexputil.Match(`charset=([\(\):\.\w-]+)`, contentType)
	if matched && strings.ToLower(contentCharset) != "utf-8" { // convert for all not utf-8 charset
		var enc encoding.Encoding
		enc, err = ianaindex.IANA.Encoding(contentCharset)
		if err != nil {
			return
		}
		bodyReader = transform.NewReader(bodyReader, enc.NewDecoder())
	}
	b, err = ioutil.ReadAll(bodyReader)
	return
}

// ReadString read string from resp.Body and close resp.Body
func ReadString(resp *http.Response) (string, error) {
	b, err := ReadAll(resp)
	if err != nil {
		return "", err
	}
	return string(b), err
}

// ReadJson read json from resp.Body and close resp.Body
func ReadJson(resp *http.Response, v interface{}) (err error) {
	defer func() {
		err = resp.Body.Close()
	}()
	err = json.NewDecoder(resp.Body).Decode(v)
	return
}

// ReadFile save file from resp.Body and close resp.Body
func ReadFile(resp *http.Response, name string) (written int64, err error) {
	var f *os.File
	f, err = fileutil.OpenWrite(name)
	if err != nil {
		return
	}
	defer func() {
		err = f.Close()
	}()
	written, err = CopyTo(resp, f)
	return
}

// CopyTo copy to dst from resp.Body and close resp.Body
func CopyTo(resp *http.Response, dst io.Writer) (written int64, err error) {
	defer func() {
		err = resp.Body.Close()
	}()
	written, err = io.Copy(dst, resp.Body)
	return
}
