package parser

import (
	"bytes"
	"encoding/json"
)

// MarshalMap
func MarshalMap(in, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(in)
	json.NewDecoder(buf).Decode(out)
}
