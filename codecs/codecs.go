package codecs

import (
	"encoding/binary"
	"errors"
	"math"
	"math/bits"
	"sync"
)

const (
	emptyString = ""
)

// errors ...
var (
	// ErrNotSupported is returned when an implement does not support
	// an operation.
	ErrNotSupported = errors.New("not supported")

	// ErrDecodeFailed is returned when decoding a byte slice into native type failed.
	ErrDecodeFailed = errors.New("failed to decode value")
)

//******************************************
// DefaultCodec Registry
//******************************************

var (
	intCodec    IntCodec
	floatCodec  FloatCodec
	textCodec   TextCodec
	bytesCodec  BytesCodec
	boolCodec   BooleanCodec
	listCodec   ListCodec
	recordCodec RecordCodec
	timeCodec   TimeCodec
)

//******************************************
// Codec Functions
//******************************************

// DecodeInt16FromBytes attempts to decode provided byte slice
// into a int16 ensuring that it has minimum length of 2.
// It uses binary.BigEndian.
func DecodeInt16FromBytes(b []byte) (int16, error) {
	de, err := DecodeUint16FromBytes(b)
	return int16(de), err
}

// DecodeUint16FromBytes attempts to decode provided byte slice
// into a uint16 ensuring that it has minimum length of 2.
// It uses binary.BigEndian.
func DecodeUint16FromBytes(b []byte) (uint16, error) {
	if len(b) < 2 {
		return 0, errors.New("byte slice length too small, must be 2")
	}

	var err error
	defer func() {
		if it := recover(); it != nil {
			err = errors.New("failed to decode byte slice with binary.BigEndian")
		}
	}()
	return binary.BigEndian.Uint16(b), err
}

// DecodeInt64FromBytes attempts to decode provided byte slice
// into a int64 ensuring that it has minimum length of 8.
// It uses binary.BigEndian.
func DecodeInt64FromBytes(b []byte) (int64, error) {
	de, err := DecodeUint64FromBytes(b)
	return int64(de), err
}

// DecodeUint64FromBytes attempts to decode provided byte slice
// into a uint64 ensuring that it has minimum length of 8.
// It uses binary.BigEndian.
func DecodeUint64FromBytes(b []byte) (uint64, error) {
	if len(b) < 8 {
		return 0, errors.New("byte slice length too small, must be 8")
	}

	var err error
	defer func() {
		if it := recover(); it != nil {
			err = errors.New("failed to decode byte slice with binary.BigEndian")
		}
	}()
	return binary.BigEndian.Uint64(b), err
}

// DecodeInt32FromBytes attempts to decode provided byte slice
// into a int32 ensuring that it has minimum length of 4.
// It uses binary.BigEndian.
func DecodeInt32FromByte(b []byte) (int32, error) {
	de, err := DecodeUint32FromBytes(b)
	return int32(de), err
}

// DecodeUint32FromBytes attempts to decode provided byte slice
// into a uint32 ensuring that it has minimum length of 4.
// It uses binary.BigEndian.
func DecodeUint32FromBytes(b []byte) (uint32, error) {
	if len(b) < 4 {
		return 0, errors.New("byte slice length too small, must be 4")
	}

	var err error
	defer func() {
		if it := recover(); it != nil {
			err = errors.New("failed to decode byte slice with binary.BigEndian")
		}
	}()
	return binary.BigEndian.Uint32(b), err
}

// EncodeInt32ToBytes encodes provided uint32 into provided
// byte ensuring byte slice has minimum of length 4.
// It uses binary.BigEndian.
func EncodeInt32ToBytes(f int32, b []byte) error {
	return EncodeUint32ToBytes(uint32(f), b)
}

// EncodeUint16ToBytes encodes provided uint16 into provided
// byte ensuring byte slice has minimum of length 2.
// It uses binary.BigEndian.
func EncodeUint16ToBytes(f uint16, b []byte) error {
	if cap(b) < 2 {
		return errors.New("required 8 length for size")
	}

	binary.BigEndian.PutUint16(b, f)
	return nil
}

// EncodeUint32ToBytes encodes provided uint32 into provided
// byte ensuring byte slice has minimum of length 4.
// It uses binary.BigEndian.
func EncodeUint32ToBytes(f uint32, b []byte) error {
	if cap(b) < 4 {
		return errors.New("required 8 length for size")
	}

	binary.BigEndian.PutUint32(b, f)
	return nil
}

// EncodeInt64ToBytes encodes provided uint64 into provided
// byte ensuring byte slice has minimum of length 8.
// It uses binary.BigEndian.
func EncodeInt64ToBytes(f int64, b []byte) error {
	return EncodeUint64ToBytes(uint64(f), b)
}

// EncodeUint64ToBytes encodes provided uint64 into provided
// byte ensuring byte slice has minimum of length 8.
// It uses binary.BigEndian.
func EncodeUint64ToBytes(f uint64, b []byte) error {
	if cap(b) < 8 {
		return errors.New("required 8 length for size")
	}

	binary.BigEndian.PutUint64(b, f)
	return nil
}

// DecodeFloat32 will decode provided uint64 value which should be in
// standard IEEE 754 binary representation, where it bit has been reversed,
// where having it's exponent appears first. It returns the float32 value.
func DecodeFloat32(f uint32) float32 {
	rbit := bits.ReverseBytes32(f)
	return math.Float32frombits(rbit)
}

// EncodeFloat64 will encode provided float value into the standard
// IEEE 754 binary representation and has it's bit reversed, having
// the exponent appearing first.
func EncodeFloat32(f float32) uint32 {
	fbit := math.Float32bits(f)
	return bits.ReverseBytes32(fbit)
}

// DecodeFloat64 will decode provided uint64 value which should be in
// standard IEEE 754 binary representation, where it bit has been reversed,
// where having it's exponent appears first. It returns the float64 value.
func DecodeFloat64(f uint64) float64 {
	rbit := bits.ReverseBytes64(f)
	return math.Float64frombits(rbit)
}

// EncodeFloat64 will encode provided float value into the standard
// IEEE 754 binary representation and has it's bit reversed, having
// the exponent appearing first.
func EncodeFloat64(f float64) uint64 {
	fbit := math.Float64bits(f)
	return bits.ReverseBytes64(fbit)
}

// EncodeVarInt32 encodes uint32 into a byte slice
// using EncodeVarInt64 after turing uint32 into uin64.
func EncodeVarInt32(x uint32) []byte {
	return EncodeVarInt64(uint64(x))
}

// EncodeUInt16 returns the encoded byte slice of a uint16 value.
func EncodeUInt16(x uint16) []byte {
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, x)
	return data
}

// EncodeVarInt64 returns the varint encoding of x.
// This is the format for the
// int32, int64, uint32, uint64, bool, and enum.
func EncodeVarInt64(x uint64) []byte {
	var buf [10]byte
	var n int
	for n = 0; x > 127; n++ {
		buf[n] = 0x80 | uint8(x&0x7F)
		x >>= 7
	}
	buf[n] = uint8(x)
	n++
	return buf[0:n]
}

// DecodeVarInt32 encodes uint32 into a byte slice
// using EncodeVarInt64 after turing uint32 into uin64.
func DecodeVarInt32(b []byte) (uint32, int) {
	v, d := DecodeVarInt64(b)
	return uint32(v), d
}

// DecodeUInt16 returns the decoded uint16 of provided byte slice which
// must be of length 2.
func DecodeUInt16(d []byte) uint16 {
	return binary.BigEndian.Uint16(d)
}

// DecodeVarInt64 reads a varint-encoded integer from the slice.
// It returns the integer and the number of bytes consumed, or
// zero if there is not enough.
// This is the format for the
// int32, int64, uint32, uint64, bool.
func DecodeVarInt64(buf []byte) (x uint64, n int) {
	for shift := uint(0); shift < 64; shift += 7 {
		if n >= len(buf) {
			return 0, 0
		}
		b := uint64(buf[n])
		n++
		x |= (b & 0x7F) << shift
		if (b & 0x80) == 0 {
			return x, n
		}
	}

	// The number is too large to represent in a 64-bit value.
	return 0, 0
}

//******************************************
// AppendWriter
//******************************************

var appendPool = sync.Pool{New: func() interface{} {
	return new(AppendWriter)
}}

// AppendWriter implements io.Writer and writes all received data into
// a underline byte slice.
type AppendWriter struct {
	C []byte
}

func (aw *AppendWriter) Write(b []byte) (int, error) {
	aw.C = append(aw.C, b...)
	return len(b), nil
}
