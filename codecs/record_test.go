package codecs_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"time"

	"github.com/influx6/faux/tests"
	"github.com/wirekit/voxa/codecs"
)

func TestRecordCodec_NativeToBinary_Basic(t *testing.T) {
	record := struct {
		Username       string   `id:"1"`
		FavoriteNumber int      `id:"2"`
		Interests      []string `id:"3"`
	}{
		FavoriteNumber: 1337,
		Username:       "bob",
		Interests:      []string{"daydreaming", "hacking"},
	}

	var codec codecs.RecordCodec
	encoded, err := codec.NativeToBinary(record, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with record codec")
	}
	tests.Passed("Should have successfully encoded value with record codec")

	if jsonEncoded, err := json.Marshal(record); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	res := &(struct {
		Username       string   `id:"1"`
		FavoriteNumber int      `id:"2"`
		Interests      []string `id:"3"`
	}{})

	err = codec.BinaryToNative(encoded, reflect.ValueOf(res))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with record codec")
	}
	tests.Passed("Should have successfully decoded value with record codec")

	if !reflect.DeepEqual(*res, record) {
		tests.Failed("Should have matching elements between input and res")
	}
	tests.Passed("Should have matching elements between input and res")
}

func TestRecordCodec_NativeToBinary_StructOnly_DuplicateTags(t *testing.T) {
	record := struct {
		Age       int    `id:"1"`
		Name      string `id:"4"`
		Address   string `id:"3"`
		IsPregant bool   `id:"4"`
	}{
		Age:       20,
		IsPregant: true,
		Name:      "bob",
		Address:   "20. Classy Street",
	}

	var codec codecs.RecordCodec
	if _, err := codec.NativeToBinary(record, []byte{}); err == nil {
		tests.Failed("Should have failed due to duplicate tag")
	}
	tests.Passed("Should have failed due to duplicate tag")
}

func TestRecordCodec_NativeToBinary_StructOnly(t *testing.T) {
	record := struct {
		Age       int    `id:"1"`
		Name      string `id:"2"`
		Address   string `id:"3"`
		IsPregant bool   `id:"4"`
	}{
		Age:       20,
		IsPregant: true,
		Name:      "bob",
		Address:   "20. Classy Street",
	}

	var codec codecs.RecordCodec
	encoded, err := codec.NativeToBinary(record, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with record codec")
	}
	tests.Passed("Should have successfully encoded value with record codec")

	if jsonEncoded, err := json.Marshal(record); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	res := &(struct {
		Age       int    `id:"1"`
		Name      string `id:"2"`
		Address   string `id:"3"`
		IsPregant bool   `id:"4"`
	}{})

	err = codec.BinaryToNative(encoded, reflect.ValueOf(res))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with record codec")
	}
	tests.Passed("Should have successfully decoded value with record codec")

	if !reflect.DeepEqual(*res, record) {
		tests.Failed("Should have matching elements between input and res")
	}
	tests.Passed("Should have matching elements between input and res")
}

func TestRecordCodec_NativeToBinary_StructWithList(t *testing.T) {
	record := struct {
		Age        int      `id:"1"`
		Name       string   `id:"2"`
		Address    string   `id:"3"`
		OtherNames []string `id:"4"`
		Bits       []byte   `id:"5"`
	}{
		Age:        20,
		Name:       "bob",
		Address:    "20. Classy Street",
		Bits:       []byte("20. Classy"),
		OtherNames: []string{"Rick Woss", "Ross Rics", "Frilino Felioi"},
	}

	var codec codecs.RecordCodec
	encoded, err := codec.NativeToBinary(record, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with record codec")
	}
	tests.Passed("Should have successfully encoded value with record codec")

	if jsonEncoded, err := json.Marshal(record); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	res := &(struct {
		Age        int      `id:"1"`
		Name       string   `id:"2"`
		Address    string   `id:"3"`
		OtherNames []string `id:"4"`
		Bits       []byte   `id:"5"`
	}{})

	err = codec.BinaryToNative(encoded, reflect.ValueOf(res))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with record codec")
	}
	tests.Passed("Should have successfully decoded value with record codec")

	if !reflect.DeepEqual(*res, record) {
		tests.Failed("Should have matching elements between input and res")
	}
	tests.Passed("Should have matching elements between input and res")
}

func TestRecordCodec_NativeToBinary_StructWithNestedList(t *testing.T) {
	record := struct {
		Age        int        `id:"1"`
		Name       string     `id:"2"`
		Address    string     `id:"3"`
		Date       time.Time  `id:"5"`
		OtherNames [][]string `id:"4"`
	}{
		Age:     20,
		Name:    "bob",
		Address: "20. Classy Street",
		Date:    time.Now(),
		OtherNames: [][]string{
			[]string{"wreckage", "went into downtown"},
			[]string{"moppers guild", "God is Love"},
			[]string{"Is His always Faithful!"},
		},
	}

	var codec codecs.RecordCodec
	encoded, err := codec.NativeToBinary(record, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with record codec")
	}
	tests.Passed("Should have successfully encoded value with record codec")

	if jsonEncoded, err := json.Marshal(record); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	res := &(struct {
		Age        int        `id:"1"`
		Name       string     `id:"2"`
		Address    string     `id:"3"`
		OtherNames [][]string `id:"4"`
		Date       time.Time  `id:"5"`
	}{})

	err = codec.BinaryToNative(encoded, reflect.ValueOf(res))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with record codec")
	}
	tests.Passed("Should have successfully decoded value with record codec")

	if res.Date.Unix() != record.Date.Unix() {
		tests.Failed("Date not matching")
	}

	if res.Name != record.Name {
		tests.Failed("Name not matching")
	}

	if res.Age != record.Age {
		tests.Failed("Age not matching")
	}

	if res.Address != record.Address {
		tests.Failed("Address not matching")
	}

	if !reflect.DeepEqual(res.OtherNames, record.OtherNames) {
		tests.Failed("Should have matching elements between input and res")
	}
	tests.Passed("Should have matching elements between input and res")
}

type Address struct {
	Value string `id:"1"`
}

func TestRecordCodec_NativeToBinary_StructWithNestedStructList(t *testing.T) {
	record := struct {
		Age        int       `id:"1"`
		Name       string    `id:"2"`
		Address    string    `id:"3"`
		Date       time.Time `id:"5"`
		OtherNames []Address `id:"4"`
	}{
		Age:     20,
		Name:    "bob",
		Address: "20. Classy Street",
		Date:    time.Now(),
		OtherNames: []Address{
			{Value: "wreckage"},
			{Value: "moppers guild"},
			{Value: "Is His always Faithful!"},
		},
	}

	var codec codecs.RecordCodec
	encoded, err := codec.NativeToBinary(record, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with record codec")
	}
	tests.Passed("Should have successfully encoded value with record codec")

	if jsonEncoded, err := json.Marshal(record); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	res := &(struct {
		Age        int       `id:"1"`
		Name       string    `id:"2"`
		Address    string    `id:"3"`
		OtherNames []Address `id:"4"`
		Date       time.Time `id:"5"`
	}{})

	err = codec.BinaryToNative(encoded, reflect.ValueOf(res))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with record codec")
	}
	tests.Passed("Should have successfully decoded value with record codec")

	if res.Date.Unix() != record.Date.Unix() {
		tests.Failed("Date not matching")
	}

	if res.Name != record.Name {
		tests.Failed("Name not matching")
	}

	if res.Age != record.Age {
		tests.Failed("Age not matching")
	}

	if res.Address != record.Address {
		tests.Failed("Address not matching")
	}

	if !reflect.DeepEqual(res.OtherNames, record.OtherNames) {
		tests.Failed("Should have matching elements between input and res")
	}
	tests.Passed("Should have matching elements between input and res")
}

func TestRecordCodec_NativeToBinary_Map(t *testing.T) {
	record := map[interface{}]interface{}{
		"age":     20,
		"name":    "bob",
		"address": "20. Classy Street",
		"date":    time.Now(),
	}

	var codec codecs.RecordCodec
	encoded, err := codec.NativeToBinary(record, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with record codec")
	}
	tests.Passed("Should have successfully encoded value with record codec")

	if jsonEncoded, err := json.Marshal(record); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	res := &(map[interface{}]interface{}{})
	err = codec.BinaryToNative(encoded, reflect.ValueOf(res))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with record codec")
	}
	tests.Passed("Should have successfully decoded value with record codec")

	if !matchAllIn(getValues(record), getValues(*res)) {
		tests.Failed("Should have matching values in input and output")
	}
	tests.Passed("Should have matching values in input and output")
}

func getValues(m map[interface{}]interface{}) []interface{} {
	var items []interface{}
	for _, val := range m {
		if tval, ok := val.(time.Time); ok {
			items = append(items, tval.Unix())
			continue
		}
		items = append(items, val)
	}
	return items
}

func matchAllIn(v, r []interface{}) bool {
	if len(v) != len(r) {
		return false
	}
	for _, k := range v {
		if !find(k, r) {
			return false
		}
	}
	return true
}

func find(v interface{}, r []interface{}) bool {
	for _, k := range r {
		if k == v {
			return true
		}
		if reflect.DeepEqual(k, v) {
			return true
		}
	}
	return false
}
