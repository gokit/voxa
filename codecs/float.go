package codecs

import (
	"errors"

	"strconv"

	"github.com/wirekit/voxa"
)

var _ voxa.Codec = FloatCodec{}

type FloatCodec struct{}

func (FloatCodec) BinaryToNative(b []byte) (interface{}, voxa.FieldID, error) {
	if len(b) < 3 {
		return nil, 0, errors.New("byte slice must be of length 2")
	}

	id := voxa.FieldID(b[1])

	val := b[2:]
	switch voxa.Atom(b[0]) {
	case voxa.Float32:
		dl, n := DecodeVarInt32(val)
		if n == 0 {
			return nil, id, ErrDecodeFailed
		}

		return DecodeFloat32(dl), id, nil
	case voxa.Float64:
		dl, n := DecodeVarInt64(val)
		if n == 0 {
			return nil, id, ErrDecodeFailed
		}

		return DecodeFloat64(dl), id, nil
	default:
		return nil, id, errors.New("byte slice must have supported type marker")
	}
}

func (FloatCodec) NativeToBinary(b interface{}, id voxa.FieldID, c []byte) ([]byte, error) {
	if val, ok := b.(float32); ok {
		enc := EncodeVarInt32(EncodeFloat32(val))
		return append(append(c, byte(voxa.Float32), byte(id)), enc...), nil
	}
	if val, ok := b.(float64); ok {
		enc := EncodeVarInt64(EncodeFloat64(val))
		return append(append(c, byte(voxa.Float64), byte(id)), enc...), nil
	}

	return nil, errors.New("type is not a float32/float64")
}

func (FloatCodec) MarshalTextualToNative(_ []byte, _ interface{}) error {
	return ErrNotSupported
}

func (FloatCodec) TextualToNative(b []byte) (interface{}, error) {
	return strconv.ParseFloat(string(b), 64)
}

func (FloatCodec) NativeToTextual(b interface{}, c []byte) ([]byte, error) {
	if val, ok := b.(float32); ok {
		return strconv.AppendFloat(c, float64(val), 'f', 10, 32), nil
	}
	if val, ok := b.(float64); ok {
		return strconv.AppendFloat(c, val, 'f', 10, 64), nil
	}
	return nil, errors.New("type is not a float32/float64")
}
