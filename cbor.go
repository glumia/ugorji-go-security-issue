package cbor

import "github.com/ugorji/go/codec"

var handle = new(codec.CborHandle)

func Unmarshal(data []byte, dst interface{}) error {
	dec := codec.NewDecoderBytes(data, handle)
	if err := dec.Decode(dst); err != nil {
		return err
	}
	return nil
}
