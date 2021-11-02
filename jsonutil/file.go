package jsonutil

import (
	"encoding/json"
	"github.com/gofika/util/fileutil"
	"os"
)

// ReadFile read struct from json file
func ReadFile(filename string, e interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(e)
	if err != nil {
		return err
	}
	return nil
}

// WriteFile write struct to json file
func WriteFile(filename string, e interface{}) error {
	f, err := fileutil.OpenWrite(filename)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	enc := json.NewEncoder(f)
	return enc.Encode(e)
}

// WriteFileIndent write struct to json file with indent
func WriteFileIndent(filename string, e interface{}, indent string) error {
	f, err := fileutil.OpenWrite(filename)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	enc := json.NewEncoder(f)
	enc.SetIndent("", indent)
	return enc.Encode(e)
}
