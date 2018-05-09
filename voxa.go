package voxa

import (
	"math"
)

const (
	// IDTagName specifies the giving name for tagging struct fields
	// with, to mark a field as matching a giving id from a struct to
	// be converted or one to be deserialized with binary stream.
	IDTagName = "id"
)

var (
	// MaxBlockCount is the maximum number of data items allowed in a single
	// block that will be decoded from a binary stream, whether when reading
	// blocks to decode an array or a map, or when reading blocks from an OCF
	// stream. This check is to ensure decoding binary data will not cause the
	// library to over allocate RAM, potentially creating a denial of service on
	// the system.
	//
	// If a particular application needs to decode binary Voxa data that
	// potentially has more data items in a single block, then this variable may
	// be modified at your discretion.
	MaxBlockCount = int64(math.MaxInt32)

	// MaxBlockSize is the maximum number of bytes that will be allocated for a
	// single block of data items when decoding from a binary stream. This check
	// is to ensure decoding binary data will not cause the library to over
	// allocate RAM, potentially creating a denial of service on the system.
	//
	// If a particular application needs to decode binary Voxa data that
	// potentially has more bytes in a single block, then this variable may be
	// modified at your discretion.
	MaxBlockSize = int64(math.MaxInt32)

	// MaxIBU8Count is the maximum number of items a []uint8/[]int8/[]byte/[]bool can contain
	// per block, with respect to the maximum allowed block size in MaxBlockSize.
	MaxIBU8Count = MaxBlockSize / 1

	// MaxIBU16Count is the maximum number of items a []uint16/[]int16 can contain
	// per block, with respect to the maximum allowed block size in MaxBlockSize.
	MaxIBU16Count = MaxBlockSize / 2

	// MaxIBU32Count is the maximum number of items a []uint32/[]int32/[]float32 can contain
	// per block, with respect to the maximum allowed block size in MaxBlockSize.
	MaxIBU32Count = MaxBlockSize / 4

	// MaxIBU64 is the maximum number of items a []uint64/[]int64/[]float64 can contain
	// per block, with respect to the maximum allowed block size in MaxBlockSize.
	MaxIBU64Count = MaxBlockSize / 8

	// MaxICU128 is the maximum number of items a []complex128 can contain
	// per block, with respect to the maximum allowed block size in MaxBlockSize.
	MaxICU128Count = MaxBlockSize / 16
)

// constants of all Type types.
const (
	Invalid Atom = iota
	Text
	Bit
	Int
	Int8
	Int16
	Int32
	Int64
	UInt
	UInt8
	UInt16
	UInt32
	UInt64
	Boolean
	Float32
	Float64
	Bytes
	List
	Record
	Time
)

// Atom is a int8 type declaration to represent different
// acceptable and convertible data types.
type Atom uint8

func (a Atom) String() string {
	switch a {
	case Invalid:
		return "invalid"
	case Text:
		return "text"
	case Int:
		return "int"
	case Int8:
		return "int8"
	case Int16:
		return "int16"
	case Int32:
		return "int32"
	case Int64:
		return "int64"
	case UInt:
		return "uint"
	case UInt8:
		return "uint8"
	case UInt16:
		return "uint16"
	case UInt32:
		return "uint32"
	case UInt64:
		return "uint64"
	case Boolean:
		return "bool"
	case Float32:
		return "float32"
	case Float64:
		return "float64"
	case Bit:
		return "bit/byte"
	case Bytes:
		return "bytes"
	case List:
		return "list"
	case Record:
		return "record"
	case Time:
		return "time"
	default:
		return "invalid"
	}
}

// FieldID sets a int8 type which is used to represent the
// ID attached to a field name.
type FieldID int8

// HeaderCodec defines a interface which exposes two methods
// to derive the representation of giving field name in a
// readable format. It is meant to have types define how
// they wish the meta to represent their associated field
// with type details.
type HeaderCodec interface {
	// FieldToBinary transforms giving field name string and it's unique id
	// into desired representation containing a type bit stored in
	// the provided slice which is returned.
	FieldToBinary(string, FieldID, Atom, []byte) ([]byte, error)

	// BinaryToField attempts to transforms provided byte slice
	// into a field name and Atom flag to both represent the
	// type desired and its type.
	BinaryToField([]byte) (string, FieldID, Atom, error)
}

// Codec defines a interface type which exposes
// conversion methods for types.
type Codec interface {
	// BinaryToNative takes giving byte slice and attempts to
	// convert binary into native.
	BinaryToNative([]byte) (interface{}, FieldID, error)

	// NativeToBinary will receive a native value, which it encodes
	// into the provided byte slice, returning provided byte slice
	// with new length.
	NativeToBinary(interface{}, FieldID, []byte) ([]byte, error)
}

// CodecTextual defines a type which include textual to native and vice-versal
// conversion methods.
type CodecTextual interface {
	Codec

	// TextualToNative takes giving byte slice containing text version
	// and attempts to convert into native.
	TextualToNative([]byte) (interface{}, error)

	// MarshalTextualToNative takes giving byte slice and attempts to
	// unmarshal/deserialize it's content into provided type. This exists
	// to allow underline implementation specify the native type to be
	// converted to from a textual representation.
	MarshalTextualToNative([]byte, interface{}) error

	// NativeToTextual takes giving native value and converts into bytes,
	// writing into provided byte slice and returns provided byte slice
	// and returning provided byte slice with new length.
	NativeToTextual(interface{}, []byte) ([]byte, error)
}
