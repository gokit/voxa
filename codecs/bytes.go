package codecs

import (
	"errors"

	"github.com/wirekit/voxa"
)

var _ voxa.Codec = BytesCodec{}

type BytesCodec struct{}

func (BytesCodec) BinaryToNative(b []byte) (interface{}, voxa.FieldID, error) {
	if len(b) < 3 {
		return nil, 0, errors.New("byte slice must be of length 2")
	}

	id := voxa.FieldID(b[1])
	if voxa.Atom(b[0]) != voxa.Bytes {
		return nil, id, errors.New("byte slice must have supported type marker")
	}

	return b[2:], id, nil
}

func (BytesCodec) NativeToBinary(b interface{}, id voxa.FieldID, c []byte) ([]byte, error) {
	if bu, ok := b.([]byte); ok {
		return append(append(c, byte(voxa.Bytes), byte(id)), bu...), nil
	}

	return nil, errors.New("type is not a []byte")
}

func (BytesCodec) MarshalTextualToNative(_ []byte, _ interface{}) error {
	return ErrNotSupported
}

func (BytesCodec) TextualToNative(b []byte) (interface{}, error) {
	return b, nil
}

func (BytesCodec) NativeToTextual(b interface{}, c []byte) ([]byte, error) {
	if bu, ok := b.([]byte); ok {
		return append(c, bu...), nil
	}
	return nil, errors.New("type is not a []byte")
}
