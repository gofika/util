package httputil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gofika/util/fileutil"
	"github.com/gofika/util/regexputil"
	"go.uber.org/multierr"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
)

// ReadAll read all data from resp and close resp.Body
func ReadAll(resp *http.Response) (b []byte, err error) {
	defer func() {
		if errClose := resp.Body.Close(); errClose != nil {
			err = multierr.Append(err, errClose)
		}
	}()
	b, err = io.ReadAll(resp.Body)
	return
}

// ReadString read string from resp.Body and close resp.Body
func ReadString(resp *http.Response) (s string, err error) {
	defer func() {
		if errClose := resp.Body.Close(); errClose != nil {
			err = multierr.Append(err, errClose)
		}
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
	b, err := io.ReadAll(bodyReader)
	s = string(b)
	return
}

// ReadJson read json from resp.Body and close resp.Body
func ReadJson(resp *http.Response, v any) (err error) {
	defer func() {
		if errClose := resp.Body.Close(); errClose != nil {
			err = multierr.Append(err, errClose)
		}
	}()
	err = json.NewDecoder(resp.Body).Decode(v)
	return
}

// SaveFile save file from resp.Body and close resp.Body
func SaveFile(resp *http.Response, name string) (written int64, err error) {
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

// CopyTo copy to dest from resp.Body and close resp.Body
func CopyTo(resp *http.Response, dest io.Writer) (written int64, err error) {
	defer func() {
		if errClose := resp.Body.Close(); errClose != nil {
			err = multierr.Append(err, errClose)
		}
	}()
	written, err = io.Copy(dest, resp.Body)
	return
}
