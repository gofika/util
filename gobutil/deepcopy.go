package gobutil

import (
	"bytes"
	"encoding/gob"
)

// DeepCopy clone object. can clone without same type
//
// Example:
//     type A struct {
//         Name string
//         Value int
//     }
//
//     type B struct {
//         Name string
//         Value int
//     }
//
//     a := &A { Name: "Jason", 100}
//     var b B
//     DeepCopy(&b, a)
//     fmt.Printf("%+v\n", b)
func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(&buf).Decode(dst)
}

