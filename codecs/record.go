package codecs

import (
	"errors"

	"reflect"

	"strconv"

	"fmt"

	"github.com/wirekit/voxa"
)

var (
	// ErrNotRecord is returned when data slice provided is not a voxa.Record type.
	ErrNotRecord = errors.New("data item is not a Record")

	// ErrMustBePointer is returned when type to be unmarshalled to is a value.
	ErrMustBePointer = errors.New("destination must be a pointer of desired type")

	// ErrTagMustBeUniqueToField is returned when a tag id number is found in another field.
	ErrTagMustBeUniqueToField = errors.New("id tag number already used")

	// ErrUnknownType is returned when type cant be processed.
	ErrUnknownType = errors.New("type unknown")

	// ErrUnknownTypeForMap is returned when type cant be made for when creating a map.
	ErrUnknownTypeForMap = errors.New("type unsupported for maps")

	// ErrTagMustBeNumber is returned when a tag contains more than digit values.
	ErrTagMustBeNumber = errors.New("id tag must contain only digits")

	// ErrTagCantBeMoreThanUint8 is returned when a tag contains more the max values for a uint8.
	ErrTagCantBeMoreThanUint8 = errors.New("id tag numbers must be less or equal to a uint8 or 255")
)

var (
	mapType       = (*map[interface{}]interface{})(nil)
	listSliceType = (*[]interface{})(nil)
	byteSliceType = (*[]byte)(nil)
	boolType      = (*bool)(nil)
	intType       = (*int)(nil)
	int8Type      = (*int8)(nil)
	int16Type     = (*int16)(nil)
	int32Type     = (*int32)(nil)
	int64Type     = (*int64)(nil)
	uintType      = (*uint)(nil)
	uint8Type     = (*uint8)(nil)
	uint16Type    = (*uint16)(nil)
	uint32Type    = (*uint32)(nil)
	uint64Type    = (*uint64)(nil)
	stringType    = (*string)(nil)
	float32Type   = (*float32)(nil)
	float64Type   = (*float64)(nil)
)

type RecordCodec struct{}

func (lc RecordCodec) BinaryToNative(b []byte, target interface{}) error {
	xl, read := DecodeVarInt64(b)
	if xl == 0 {
		return ErrInvalidNoSize
	}

	headerFrame := b[:read+2]
	dataFrame := b[read+2:]

	if len(dataFrame) < int(xl-2) {
		return ErrInvalidDataSlice
	}

	//itemCount := lc.countBinaryItems(dataFrame)
	//fmt.Printf("BinaryToNative:Items %d\n", itemCount)

	header := headerFrame[read:]
	htype := voxa.Atom(header[0])
	if htype != voxa.Record {
		return ErrNotRecord
	}

	var itemVal reflect.Value
	if itval, ok := target.(reflect.Value); ok {
		itemVal = itval
	} else {
		itemVal = reflect.ValueOf(target)
	}

	// It's important the supplid type is a pointer else we will face
	// error with CanSet for any of it's field.
	if itemVal.Kind() != reflect.Ptr {
		return ErrMustBePointer
	}

	itemVal = itemVal.Elem()
	item := itemVal.Type()
	if item.Kind() == reflect.Ptr {
		item = item.Elem()
	}

	if item.Kind() != reflect.Struct && item.Kind() != reflect.Map {
		return errors.New("only struct and map types acceptable")
	}

	return lc.binaryToNativeWithParent(dataFrame, itemVal, item)
}

func (lc RecordCodec) binaryToNativeWithParent(dataFrame []byte, parent reflect.Value, pType reflect.Type) error {
	var err error
	for len(dataFrame) > 0 {
		subXL, subRead := DecodeVarInt64(dataFrame)
		if subXL == 0 {
			return ErrInvalidNoSize
		}

		totalFrame := subRead + int(subXL)
		frame := dataFrame[0:totalFrame]
		subDataFrame := frame[subRead:]

		// if giving field is not found, maybe type does not has corresponding
		// destination, so skip.
		hid := subDataFrame[1]
		hidText := strconv.Itoa(int(hid))

		// we are dealing with a sublist, then we must backtrack
		// and ensure to have full header and body.
		atom := voxa.Atom(subDataFrame[0])

		var total int
		if atom == voxa.Record || atom == voxa.List {
			total = countBinaryItems(subDataFrame[2:])
			subDataFrame = frame
		}

		var field reflect.StructField

		if parent.Kind() == reflect.Struct {
			field, err = getFieldByTagAndValue(pType, voxa.IDTagName, hidText)
			if err != nil {
				// Reduce current length of slice.
				dataFrame = dataFrame[totalFrame:]
				continue
			}
		}

		// get new version value field.
		if err := lc.binaryToNativeItem(subDataFrame, int(hid), total, atom, parent, field); err != nil && err != ErrSkipErr {
			return err
		}

		// Reduce current length of slice.
		dataFrame = dataFrame[totalFrame:]
	}

	return nil
}

// BinaryToNativeItem attempts to convert a singular item within provided data within the byte slice
// into the provided reflect.Value, it ensures the internal data field ID attached in the encoded
// data match the provided. It returns an error if the id does not match, or if the value could not
// be decoded safely, more so, the value must be settable.
func (lc RecordCodec) binaryToNativeItem(data []byte, pos int, count int, atom voxa.Atom, parent reflect.Value, field reflect.StructField) error {
	var dest reflect.Value

	if field.Type != nil {
		if atom != voxa.List {
			dest = reflect.New(field.Type)
			if dest.Kind() == reflect.Ptr {
				dest = dest.Elem()
			}

			if !dest.CanSet() {
				return ErrValueUnsettable
			}
		} else {
			dest = reflect.MakeSlice(field.Type, 0, count)
		}
	} else {
		switch atom {
		case voxa.List:
			target := make([]interface{}, count)
			dest = reflect.ValueOf(target)
		case voxa.Record:
			dest = reflect.ValueOf([]map[interface{}]interface{}{})
		}
	}

	switch atom {
	case voxa.List:
		resDest, err := listCodec.BinaryToNative(data, dest)
		if err != nil {
			return err
		}

		newDest := reflect.ValueOf(resDest)
		if newDest.Kind() == reflect.Ptr && dest.Kind() != reflect.Ptr {
			dest = newDest.Elem()
		} else {
			dest = newDest
		}
	case voxa.Record:
		if err := recordCodec.BinaryToNative(data, dest); err != nil {
			return err
		}
	case voxa.Text:
		value, _, err := textCodec.BinaryToNative(data)
		if err != nil {
			return err
		}

		dest = reflect.ValueOf(value)
	case voxa.Time:
		value, _, err := timeCodec.BinaryToNative(data)
		if err != nil {
			return err
		}

		dest = reflect.ValueOf(value)
	case voxa.Bytes:
		value, _, err := bytesCodec.BinaryToNative(data)
		if err != nil {
			return err
		}
		dest = reflect.ValueOf(value)
	case voxa.Boolean:
		value, _, err := boolCodec.BinaryToNative(data)
		if err != nil {
			return err
		}
		dest = reflect.ValueOf(value)
	case voxa.Float64, voxa.Float32:
		value, _, err := floatCodec.BinaryToNative(data)
		if err != nil {
			return err
		}
		dest = reflect.ValueOf(value)
	case voxa.Int, voxa.UInt, voxa.UInt8, voxa.UInt16, voxa.UInt32, voxa.UInt64,
		voxa.Int8, voxa.Int16, voxa.Int32, voxa.Int64:
		value, _, err := intCodec.BinaryToNative(data)
		if err != nil {
			return err
		}

		dest = reflect.ValueOf(value)
	default:
		return errors.New("unknown type")
	}

	switch parent.Kind() {
	case reflect.Struct:
		ff := parent.FieldByName(field.Name)

		if ff.Kind() != reflect.Ptr && dest.Kind() == reflect.Ptr {
			dest = dest.Elem()
		}

		ff.Set(dest)
	case reflect.Map:
		parent.SetMapIndex(reflect.ValueOf(pos), dest)
	}

	return nil
}

func (lc RecordCodec) NativeToBinary(b interface{}, c []byte) ([]byte, error) {
	return lc.NativeToBinaryFrom(b, 0, c)
}

func (lc RecordCodec) NativeToBinaryFrom(b interface{}, id voxa.FieldID, c []byte) ([]byte, error) {
	item := reflect.ValueOf(b)
	if item.Kind() == reflect.Ptr {
		item = item.Elem()
	}

	if item.Kind() != reflect.Struct && item.Kind() != reflect.Map {
		return nil, errors.New("only map and struct types acceptable")
	}

	itemTypeElem := item.Type()
	itemTypeSizeUPtr := int(itemTypeElem.Size())
	itemTypeConservativeSize := itemTypeSizeUPtr

	buffer := slicePool.Get(itemTypeConservativeSize)
	defer buffer.Discard()

	var err error
	base := buffer.Data[:0]

	switch item.Kind() {
	case reflect.Map:
		for id, key := range item.MapKeys() {
			indexValue := item.MapIndex(key)
			base, err = nativeItemToBinary(indexValue.Interface(), voxa.FieldID(id+1), base)
			if err != nil && err != ErrSkipErr {
				return c, err
			}
		}
	case reflect.Struct:
		totalElements := item.NumField()
		seen := map[uint64]bool{}
		for i := 0; i < totalElements; i++ {
			indexValue := item.Field(i)
			indexType := itemTypeElem.Field(i)
			tag := indexType.Tag.Get(voxa.IDTagName)

			// if tag is a dash then skip field.
			if tag == "-" {
				continue
			}

			if tag == "" {
				return c, fmt.Errorf("field %q for %q requires a 'id' tag", indexType.Name, item.String())
			}

			tagValue, err := strconv.ParseUint(tag, 10, 8)
			if err != nil {
				if err == strconv.ErrRange {
					return c, ErrTagCantBeMoreThanUint8
				}

				return c, ErrTagMustBeNumber
			}

			if seen[tagValue] {
				return c, ErrTagMustBeUniqueToField
			}

			seen[tagValue] = true

			base, err = nativeItemToBinary(indexValue.Interface(), voxa.FieldID(tagValue), base)
			if err != nil && err != ErrSkipErr {
				return c, err
			}
		}
	}

	// payload is the total length of encoded contents + length of items count slice + 2 bits for flags.
	c = append(c, EncodeVarInt64(uint64(len(base)+2))...)
	c = append(c, byte(voxa.Record), byte(id))
	c = append(c, base...)
	return c, nil
}

func getFieldByTagAndValue(tl reflect.Type, tag string, value string) (reflect.StructField, error) {
	if tl.Kind() == reflect.Ptr {
		tl = tl.Elem()
	}

	for i := 0; i < tl.NumField(); i++ {
		field := tl.Field(i)
		if field.Tag.Get(tag) == value {
			return field, nil
		}
	}

	return reflect.StructField{}, errors.New("not found")
}
