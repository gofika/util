package gobutil

import (
	"encoding/gob"
	"os"

	"github.com/gofika/util/fileutil"
)

// ReadFile read struct from gob stream file
func ReadFile(filename string, e any) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(f)
	err = dec.Decode(e)
	if err != nil {
		return err
	}
	return nil
}

// WriteFile write struct to gob stream file
func WriteFile(filename string, e any) error {
	f, err := fileutil.OpenWrite(filename)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	enc := gob.NewEncoder(f)
	return enc.Encode(e)
}
