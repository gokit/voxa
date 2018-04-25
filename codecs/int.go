package codecs

import (
	"errors"

	"strconv"

	"math"

	"github.com/wirekit/voxa"
)

var _ voxa.Codec = IntCodec{}

type IntCodec struct{}

func (IntCodec) BinaryToNative(b []byte) (interface{}, voxa.FieldID, error) {
	if len(b) < 3 {
		return nil, 0, errors.New("byte slice must be of length 2")
	}

	id := voxa.FieldID(b[1])
	val := b[2:]
	switch voxa.Atom(b[0]) {
	case voxa.Int:
		dl, n := DecodeVarInt64(val)
		if n == 0 {
			return nil, id, ErrDecodeFailed
		}

		return int(dl), id, nil
	case voxa.UInt:
		dl, n := DecodeVarInt64(val)
		if n == 0 {
			return nil, id, ErrDecodeFailed
		}

		return uint(dl), id, nil
	case voxa.UInt8:
		return uint8(val[0]), id, nil
	case voxa.Int8:
		return int8(val[0]), id, nil
	case voxa.Int16:
		dl, err := DecodeUint16FromBytes(val)
		if err != nil {
			return nil, id, err
		}

		return int16(dl), id, nil
	case voxa.UInt16:
		dl, err := DecodeUint16FromBytes(val)
		if err != nil {
			return nil, id, err
		}

		return uint16(dl), id, nil
	case voxa.Int32:
		dl, n := DecodeVarInt32(val)
		if n == 0 {
			return nil, id, ErrDecodeFailed
		}

		return int32(dl), id, nil
	case voxa.UInt32:
		dl, n := DecodeVarInt32(val)
		if n == 0 {
			return nil, id, ErrDecodeFailed
		}

		return uint32(dl), id, nil
	case voxa.Int64:
		dl, n := DecodeVarInt64(val)
		if n == 0 {
			return nil, id, ErrDecodeFailed
		}

		return int64(dl), id, nil
	case voxa.UInt64:
		dl, n := DecodeVarInt64(val)
		if n == 0 {
			return nil, id, ErrDecodeFailed
		}

		return dl, id, nil
	default:
		return nil, id, errors.New("byte slice must have supported type marker")
	}
}

func (IntCodec) NativeToBinary(b interface{}, id voxa.FieldID, c []byte) ([]byte, error) {
	switch val := b.(type) {
	case uint:
		if val < math.MaxUint32 {
			return append(append(c, byte(voxa.UInt), byte(id)), EncodeVarInt32(uint32(val))...), nil
		} else {
			return append(append(c, byte(voxa.UInt), byte(id)), EncodeVarInt64(uint64(val))...), nil
		}
	case uint8:
		return append(c, byte(voxa.UInt8), byte(id), val), nil
	case uint16:
		return append(append(c, byte(voxa.UInt16), byte(id)), EncodeUInt16(val)...), nil
	case uint32:
		return append(append(c, byte(voxa.UInt32), byte(id)), EncodeVarInt32(val)...), nil
	case uint64:
		return append(append(c, byte(voxa.UInt64), byte(id)), EncodeVarInt64(val)...), nil
	case int:
		if val < math.MaxInt32 {
			return append(append(c, byte(voxa.Int), byte(id)), EncodeVarInt32(uint32(val))...), nil
		} else {
			return append(append(c, byte(voxa.Int), byte(id)), EncodeVarInt64(uint64(val))...), nil
		}
	case int8:
		return append(c, byte(voxa.Int8), byte(id), uint8(val)), nil
	case int16:
		return append(append(c, byte(voxa.Int16), byte(id)), EncodeUInt16(uint16(val))...), nil
	case int32:
		return append(append(c, byte(voxa.Int32), byte(id)), EncodeVarInt32(uint32(val))...), nil
	case int64:
		return append(append(c, byte(voxa.Int64), byte(id)), EncodeVarInt64(uint64(val))...), nil
	}

	return nil, errors.New("type is not a int/uint")
}

func (IntCodec) MarshalTextualToNative(_ []byte, _ interface{}) error {
	return ErrNotSupported
}

func (IntCodec) TextualToNative(b []byte) (interface{}, error) {
	return strconv.ParseInt(string(b), 10, 64)
}

func (IntCodec) NativeToTextual(b interface{}, c []byte) ([]byte, error) {
	switch val := b.(type) {
	case uint:
		return []byte(strconv.FormatUint(uint64(val), 10)), nil
	case uint8:
		return []byte(strconv.FormatUint(uint64(val), 10)), nil
	case uint16:
		return []byte(strconv.FormatUint(uint64(val), 10)), nil
	case uint32:
		return []byte(strconv.FormatUint(uint64(val), 10)), nil
	case uint64:
		return []byte(strconv.FormatUint(val, 10)), nil
	case int8:
		return []byte(strconv.FormatInt(int64(val), 10)), nil
	case int16:
		return []byte(strconv.FormatInt(int64(val), 10)), nil
	case int32:
		return []byte(strconv.FormatInt(int64(val), 10)), nil
	case int:
		return []byte(strconv.FormatInt(int64(val), 10)), nil
	case int64:
		return []byte(strconv.FormatInt(val, 10)), nil
	}

	return nil, errors.New("type is not a int/uint")
}
