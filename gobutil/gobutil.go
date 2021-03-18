package gobutil

import (
	"encoding/gob"
	"github.com/leaker/util/fileutil"
	"os"
)

// LoadFile load gob stream from file
func LoadFile(filename string, e interface{}) error {
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

// SaveFile save gob stream to file
func SaveFile(filename string, e interface{}) (err error) {
	fileutil.EnsureDirExists(filename)
	var f *os.File
	f, err = os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()
	enc := gob.NewEncoder(f)
	return enc.Encode(e)
}
