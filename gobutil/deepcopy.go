package gobutil

import (
	"bytes"
	"encoding/gob"
)

// DeepCopy clone object. can clone without same type
//
// Example:
//
//	type Foo struct {
//	    Name string
//	    Value int
//	}
//
//	type Bar struct {
//	    Name string
//	    Value int
//	}
//
//	foo := &Foo { "Jason", 100}
//	var bar Bar
//	DeepCopy(&bar, foo)
//	fmt.Printf("%+v\n", bar)
func DeepCopy(dst, src any) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(&buf).Decode(dst)
}
