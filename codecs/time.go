package codecs

import (
	"errors"
	"strconv"

	"time"

	"github.com/wirekit/voxa"
)

var _ voxa.Codec = TimeCodec{}

type TimeCodec struct{}

func (TimeCodec) BinaryToNative(b []byte) (interface{}, voxa.FieldID, error) {
	if len(b) < 3 {
		return nil, 0, errors.New("byte slice must be of length 2")
	}

	id := voxa.FieldID(b[1])
	if voxa.Atom(b[0]) != voxa.Time {
		return nil, id, errors.New("byte slice must have supported type marker")
	}

	tick, err := time.Parse(time.RFC3339, string(b[2:]))
	return tick, id, err
}

func (TimeCodec) NativeToBinary(b interface{}, id voxa.FieldID, c []byte) ([]byte, error) {
	if val, ok := b.(time.Time); ok {
		formatted := val.Format(time.RFC3339)
		return append(append(c, byte(voxa.Time), byte(id)), formatted...), nil
	}
	return nil, errors.New("only string type supported")
}

func (TimeCodec) MarshalTextualToNative(_ []byte, _ interface{}) error {
	return ErrNotSupported
}

func (TimeCodec) TextualToNative(b []byte) (interface{}, error) {
	return strconv.Unquote(string(b))
}

func (TimeCodec) NativeToTextual(b interface{}, c []byte) ([]byte, error) {
	if value, ok := b.(string); ok {
		return append(c, strconv.Quote(value)...), nil
	}
	return nil, errors.New("only string type supported")
}
