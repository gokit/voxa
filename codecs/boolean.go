package codecs

import (
	"bytes"
	"errors"

	"github.com/wirekit/voxa"
)

var (
	_          voxa.Codec = BooleanCodec{}
	trueBytes             = []byte("true")
	falseBytes            = []byte("false")
)

const (
	on  = byte(1)
	off = byte(0)
)

type BooleanCodec struct{}

func (BooleanCodec) BinaryToNative(b []byte) (interface{}, voxa.FieldID, error) {
	if len(b) != 3 {
		return nil, 0, errors.New("byte slice must be of length 2")
	}

	id := voxa.FieldID(b[1])
	if voxa.Atom(b[0]) != voxa.Boolean {
		return nil, id, errors.New("byte slice must have supported type marker")
	}

	switch b[2] {
	case off:
		return false, id, nil
	case on:
		return true, id, nil
	}

	return nil, id, errors.New("bytes slice must either be 1 or 0")
}

func (BooleanCodec) NativeToBinary(b interface{}, f voxa.FieldID, c []byte) ([]byte, error) {
	if flag, ok := b.(bool); ok {
		if flag {
			return append(c, byte(voxa.Boolean), byte(f), on), nil
		}
		return append(c, byte(voxa.Boolean), byte(f), off), nil
	}

	return nil, errors.New("type is not a bool/boolean")
}

func (BooleanCodec) MarshalTextualToNative(_ []byte, _ interface{}) error {
	return ErrNotSupported
}

func (BooleanCodec) TextualToNative(b []byte) (interface{}, error) {
	area := len(b)
	switch area {
	case 4:
		if bytes.Equal(b[:4], trueBytes) {
			return true, nil
		}
		return nil, errors.New("text contents must be exactly 'true' or 'false'")
	case 5:
		if bytes.Equal(b[:5], falseBytes) {
			return false, nil
		}
		return nil, errors.New("text contents must be exactly 'true' or 'false'")
	default:
		return nil, errors.New("bytes for bool must either be 4 -r 5 in length")
	}
}

func (BooleanCodec) NativeToTextual(b interface{}, c []byte) ([]byte, error) {
	if flag, ok := b.(bool); ok {
		if flag {
			return append(c, "true"...), nil
		}
		return append(c, "false"...), nil
	}
	return nil, errors.New("type is not a bool/boolean")
}
