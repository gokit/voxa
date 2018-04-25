package codecs

import (
	"errors"
	"strings"

	"github.com/wirekit/voxa"
)

var (
	_ voxa.HeaderCodec = HeaderCodec{}
)

// HeaderCodec implements the voxa.HeaderCodec providing
// methods to turn giving field names and associated data into
// a byte slice with format: `[FieldID][Atom][Name Bytes...]`
// and vice-versa.
type HeaderCodec struct{}

func (HeaderCodec) BinaryToField(b []byte) (string, voxa.FieldID, voxa.Atom, error) {
	if len(b) <= 2 {
		return emptyString, 0, voxa.Invalid, errors.New("field byte slice must be longer than 2")
	}

	tp := voxa.Atom(b[0])
	if tp > voxa.Record {
		return emptyString, 0, voxa.Invalid, errors.New("field byte slice must have type bit within supported")
	}

	id := voxa.FieldID(b[1])
	return string(b[2:]), id, tp, nil
}

func (HeaderCodec) FieldToBinary(name string, id voxa.FieldID, ty voxa.Atom, b []byte) ([]byte, error) {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return nil, errors.New("field name must not be an empty string")
	}

	if ty > voxa.Record {
		return nil, errors.New("field name type bit is not supported")
	}

	b = append(b, byte(id), byte(ty))
	return append(b, name...), nil
}
