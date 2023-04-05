package yamlutil

import (
	"os"

	"github.com/gofika/util/fileutil"
	"gopkg.in/yaml.v3"
)

// ReadFile read struct from yml file
func ReadFile(filename string, e any) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	dec := yaml.NewDecoder(f)
	err = dec.Decode(e)
	if err != nil {
		return err
	}
	return nil
}

// WriteFile write struct to yml file
func WriteFile(filename string, e any) error {
	f, err := fileutil.OpenWrite(filename)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	enc := yaml.NewEncoder(f)
	return enc.Encode(e)
}

// WriteFileIndent write struct to yml file with indent
func WriteFileIndent(filename string, e any, spaces int) error {
	f, err := fileutil.OpenWrite(filename)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	enc := yaml.NewEncoder(f)
	enc.SetIndent(spaces)
	return enc.Encode(e)
}
