package cbor

import (
	"testing"
)

func TestOutOfMem1(t *testing.T) {
	var f []byte
	Unmarshal([]byte("\x9b\x00\x00000000"), f)
}

func TestOutOfMem2(t *testing.T) {
	var f []byte
	Unmarshal([]byte("\x9b\x00\x00\x81112233"), f)
}
