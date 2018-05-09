Voxa
------------
[![Go Report Card](https://goreportcard.com/badge/github.com/wirekit/voxa)](https://goreportcard.com/report/github.com/wirekit/voxa)
[![Travis CI](https://travis-ci.org/wirekit/voxa.svg?master=branch)](https://travis-ci.org/wirekit/voxa)

Voxa is a binary-compact message format suitable for delivery Go types over the wire with minimal memory usage.
It removes all meta-data and encodes into a binary format where a struct fields are simply marked by a id value, this
makes it highly unsuitable for map types.

Voxa uses `id` tags as the means of identifying fields to be encoded and fields which would receive said encoding, where
associated types must match.

## Supported Types

Voxa has working support for most Go types as listed below:

- bool
- string
- int/uint
- int8/uint8/byte
- int16/uint16
- int32/uint32
- int64/uint64
- float32/float64
- Struct
- []byte
- []{string, uint8/16/32/64, int8/16/32/64, float32/64, Struct}

In voxa, `Maps` are special in time, they do not contain any meta-data like structs about the fields, hence when
voxa decodes an encoded map, it uses the id values has keys. Hence its requires more work to get such information properly, which
makes the use of struct's more suitable.

## Install

```bash
go get -u github.com/wirekit/voxa
```

## Example

```go
record := struct {
    Age        int      `id:"1"`
    Name       string   `id:"2"`
    Address    string   `id:"3"`
    OtherNames []string `id:"4"`
}{
    Age:        20,
    Name:       "bob",
    Address:    "20. Classy Street",
    OtherNames: []string{"Rick Woss", "Ross Rics", "Frilino Felioi"},
}

var codec codecs.RecordCodec
encoded, err := codec.NativeToBinary(record, []byte{})
if err != nil {
    log.Fatal(err)
}

res := &(struct {
    Age        int      `id:"1"`
    Name       string   `id:"2"`
    Address    string   `id:"3"`
    OtherNames []string `id:"4"`
}{})

err = codec.BinaryToNative(encoded, reflect.ValueOf(res))
if err != nil {
    log.Fatal(err)
}

if !reflect.DeepEqual(*res, record) {
    log.Fatal("not matching")
}

```