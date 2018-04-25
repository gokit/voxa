Voxa
------------
Voxa is a binary-compact message format suitable for delivery Go types over the wire with minimal memory usage. It provides
a schemaless binary format where a struct details the desired data to be delivered over the wire.

Voxa uses `id` tags as the means of identifying fields to be encoded and fields which would receive said encoding, where
associated types must match.


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