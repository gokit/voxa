package codecs

import (
	"errors"

	"reflect"

	"time"

	"github.com/influx6/faux/pools/pbytes"
	"github.com/wirekit/voxa"
)

// errors ...
var (
	// ErrValueUnsettable is returned when the reflect.Value is not settable.
	ErrValueUnsettable = errors.New("value can not be set on reflect type")

	// ErrInvalidNoSize is returned when provided data slice fails to match length standard.
	ErrInvalidNoSize = errors.New("invalid data, byte slice must have length length")

	// ErrInvalidDataSlice is returned when provided data slice fails to match length standard.
	ErrInvalidDataSlice = errors.New("invalid data, byte slice not matching expected data frame length")

	// ErrNotList is returned when data slice provided is not a voxa.List type.
	ErrNotList = errors.New("data item is not a list")

	// ErrInvalidFieldID is returned when the field id does not match expected.
	ErrInvalidFieldID = errors.New("data id does not match expected field id")

	// ErrSkipErr is returned when codec is not available for a giving type.
	ErrSkipErr = errors.New("no codec available for type")
)

var (
	slicePool = pbytes.NewBitsPool(32, 100)
)

type ListCodec struct{}

func (lc ListCodec) BinaryToNative(b []byte, target interface{}) (interface{}, error) {
	xl, read := DecodeVarInt64(b)
	if xl == 0 {
		return nil, ErrInvalidNoSize
	}

	headerFrame := b[:read+2]
	dataFrame := b[read+2:]

	if len(dataFrame) < int(xl-2) {
		return nil, ErrInvalidDataSlice
	}

	//itemCount := countBinaryItems(dataFrame)
	//fmt.Printf("BinaryToNative:Items %d\n", itemCount)

	header := headerFrame[read:]
	htype := voxa.Atom(header[0])
	if htype != voxa.List {
		return nil, ErrNotList
	}

	var itemVal reflect.Value
	if itval, ok := target.(reflect.Value); ok {
		itemVal = itval
	} else {
		itemVal = reflect.ValueOf(target)
	}

	item := itemVal.Type()
	if item.Kind() == reflect.Ptr {
		item = item.Elem()
	}

	if item.Kind() != reflect.Slice {
		return nil, errors.New("only array and slice types acceptable")
	}

	typeKind := item.Elem()

	var err error
	for len(dataFrame) > 0 {
		subXL, subRead := DecodeVarInt64(dataFrame)
		if subXL == 0 {
			return nil, ErrInvalidNoSize
		}

		totalFrame := subRead + int(subXL)
		frame := dataFrame[0:totalFrame]
		subDataFrame := frame[subRead:]

		var newValue reflect.Value
		// we are dealing with a sublist, then we must backtrack
		// and ensure to have full header and body.
		atom := voxa.Atom(subDataFrame[0])
		switch atom {
		case voxa.List:
			itemCount := countBinaryItems(subDataFrame[2:])
			newValue = reflect.MakeSlice(typeKind, 0, int(itemCount))
			subDataFrame = frame
		case voxa.Record:
			newValue = reflect.New(typeKind)
			subDataFrame = frame
		default:
			newValue = reflect.New(typeKind).Elem()
		}

		newValue, err = lc.binaryToNativeItem(atom, subDataFrame, newValue)
		if err != nil && err != ErrSkipErr {
			return nil, err
		}

		if typeKind.Kind() != reflect.Ptr && newValue.Kind() == reflect.Ptr {
			newValue = newValue.Elem()
		}

		itemVal = reflect.Append(itemVal, newValue)

		// Reduce current length of slice.
		dataFrame = dataFrame[totalFrame:]
	}

	return itemVal.Interface(), nil
}

// BinaryToNativeItem attempts to convert a singular item within provided data within the byte slice
// into the provided reflect.Value, it ensures the internal data field ID attached in the encoded
// data match the provided. It returns an error if the id does not match, or if the value could not
// be decoded safely, more so, the value must be settable.
func (lc ListCodec) binaryToNativeItem(atom voxa.Atom, data []byte, dest reflect.Value) (reflect.Value, error) {
	if atom < voxa.List && !dest.CanSet() {
		return dest, ErrValueUnsettable
	}

	switch atom {
	case voxa.Record:
		err := recordCodec.BinaryToNative(data, dest)
		if err != nil {
			return dest, err
		}
	case voxa.List:
		value, err := listCodec.BinaryToNative(data, dest)
		if err != nil {
			return dest, err
		}

		dest = reflect.ValueOf(value)
	case voxa.Time:
		value, _, err := timeCodec.BinaryToNative(data)
		if err != nil {
			return dest, err
		}
		dest.Set(reflect.ValueOf(value))
	case voxa.Text:
		value, _, err := textCodec.BinaryToNative(data)
		if err != nil {
			return dest, err
		}
		dest.Set(reflect.ValueOf(value))
	case voxa.Bytes:
		value, _, err := bytesCodec.BinaryToNative(data)
		if err != nil {
			return dest, err
		}
		dest.Set(reflect.ValueOf(value))
	case voxa.Boolean:
		value, _, err := boolCodec.BinaryToNative(data)
		if err != nil {
			return dest, err
		}
		dest.Set(reflect.ValueOf(value))
	case voxa.Float64, voxa.Float32:
		value, _, err := floatCodec.BinaryToNative(data)
		if err != nil {
			return dest, err
		}
		dest.Set(reflect.ValueOf(value))
	case voxa.Int, voxa.UInt, voxa.UInt8, voxa.UInt16, voxa.UInt32, voxa.UInt64,
		voxa.Int8, voxa.Int16, voxa.Int32, voxa.Int64:

		value, _, err := intCodec.BinaryToNative(data)
		if err != nil {
			return dest, err
		}

		dest.Set(reflect.ValueOf(value))
	}

	return dest, nil
}

func (lc ListCodec) NativeToBinary(b interface{}, c []byte) ([]byte, error) {
	return lc.NativeToBinaryFrom(b, 0, c)
}

func (lc ListCodec) NativeToBinaryFrom(b interface{}, id voxa.FieldID, c []byte) ([]byte, error) {
	item := reflect.ValueOf(b)
	if item.Kind() == reflect.Ptr {
		item = item.Elem()
	}

	if item.Kind() != reflect.Array && item.Kind() != reflect.Slice {
		return nil, errors.New("only array and slice types acceptable")
	}

	totalElements := item.Len()

	itemType := item.Type()
	itemTypeElem := itemType.Elem()
	itemTypeSizeUPtr := int(itemTypeElem.Size())
	itemTypeConservativeSize := itemTypeSizeUPtr * totalElements

	buffer := slicePool.Get(itemTypeConservativeSize)
	defer buffer.Discard()

	var err error
	base := buffer.Data[:0]

	for i := 0; i < totalElements; i++ {
		indexValue := item.Index(i)
		base, err = nativeItemToBinary(indexValue.Interface(), voxa.FieldID(i), base)
		if err != nil {
			return c, err
		}
	}

	// payload is the total length of encoded contents + length of items count slice + 2 bits for flags.
	c = append(c, EncodeVarInt64(uint64(len(base)+2))...)
	c = append(c, byte(voxa.List), byte(id))
	c = append(c, base...)
	return c, nil
}

// NativeItemToBinary attempts to convert a non-slice/non-array type but basic data types, structs
// which are elements of a slice into it's basic type as part of a list identified with a Field id.
func nativeItemToBinary(b interface{}, id voxa.FieldID, c []byte) ([]byte, error) {
	itemType := reflect.TypeOf(b)
	if itemType.Kind() == reflect.Ptr {
		itemType = itemType.Elem()
	}

	switch itemType.Kind() {
	case reflect.Bool:
		content := slicePool.Get(32)
		defer content.Discard()

		encoded, err := boolCodec.NativeToBinary(b, id, content.Data[:0])
		if err != nil {
			return nil, err
		}

		// Calculate total length of encoded value + length of type flag (1).
		return append(append(c, EncodeVarInt64(uint64(len(encoded)))...), encoded...), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// We can guarantee we only need 3 bits for storing a bool
		// in voxa format.
		content := slicePool.Get(32)
		defer content.Discard()

		encoded, err := intCodec.NativeToBinary(b, id, content.Data[:0])
		if err != nil {
			return nil, err
		}

		// Calculate total length of encoded value + length of type flag (1).
		return append(append(c, EncodeVarInt64(uint64(len(encoded)))...), encoded...), nil
	case reflect.String:
		textLength := len(b.(string)) + 5
		content := slicePool.Get(textLength)
		defer content.Discard()

		encoded, err := textCodec.NativeToBinary(b, id, content.Data[:0])
		if err != nil {
			return nil, err
		}

		// Calculate total length of encoded value + length of type flag (1).
		return append(append(c, EncodeVarInt64(uint64(len(encoded)))...), encoded...), nil
	case reflect.Float32, reflect.Float64:
		// We can guarantee there will enough space to cover
		// this write has the slicePool has a 32bit start length.
		content := slicePool.Get(32)
		defer content.Discard()

		encoded, err := floatCodec.NativeToBinary(b, id, content.Data[:0])
		if err != nil {
			return nil, err
		}

		// Calculate total length of encoded value + length of type flag (1).
		return append(append(c, EncodeVarInt64(uint64(len(encoded)))...), encoded...), nil
	case reflect.Struct, reflect.Map:
		// We can guarantee there will enough space to cover
		// this write has the slicePool has a 32bit start length.
		content := slicePool.Get(1024)
		defer content.Discard()

		if _, ok := b.(time.Time); ok {
			encoded, err := timeCodec.NativeToBinary(b, id, content.Data[:0])
			if err != nil {
				return nil, err
			}

			// Calculate total length of encoded value + length of type flag (1).
			return append(append(c, EncodeVarInt64(uint64(len(encoded)))...), encoded...), nil
		}

		encoded, err := recordCodec.NativeToBinaryFrom(b, id, content.Data[:0])
		if err != nil {
			return nil, err
		}

		// Calculate total length of encoded value + length of type flag (1).
		return append(c, encoded...), nil
	case reflect.Slice:
		// We can guarantee there will enough space to cover
		// this write has the slicePool has a 32bit start length.
		content := slicePool.Get(1024)
		defer content.Discard()

		encoded, err := listCodec.NativeToBinaryFrom(b, id, content.Data[:0])
		if err != nil {
			return nil, err
		}

		// Calculate total length of encoded value + length of type flag (1).
		return append(c, encoded...), nil
	}

	return nil, ErrSkipErr
}

func countBinaryItems(b []byte) int {
	var seen int
	for len(b) > 0 {
		subarea, read := DecodeVarInt64(b)
		if read == 0 {
			return seen
		}

		frame := int(subarea) + read
		b = b[frame:]
		seen++
	}
	return seen
}
