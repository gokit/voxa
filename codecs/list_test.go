package codecs_test

import (
	"testing"

	"reflect"

	"encoding/json"

	"github.com/influx6/faux/tests"
	"github.com/wirekit/voxa/codecs"
)

func TestListCodec_NativeToBinary_UInt(t *testing.T) {
	contents := []uint{1, 2, 3, 4}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([]uint{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([]uint)
	if !ok {
		tests.Failed("Should have received slice of uint []uint")
	}
	tests.Passed("Should have received slice of uint []uint")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}

func TestListCodec_NativeToBinary_UInt8(t *testing.T) {
	contents := []uint8{1, 2, 3, 4}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([]uint8{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([]uint8)
	if !ok {
		tests.Failed("Should have received slice of uint []uint")
	}
	tests.Passed("Should have received slice of uint []uint")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}

func TestListCodec_NativeToBinary_UInt16(t *testing.T) {
	contents := []uint16{1, 2, 3, 4}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([]uint16{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([]uint16)
	if !ok {
		tests.Failed("Should have received slice of uint []uint")
	}
	tests.Passed("Should have received slice of uint []uint")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}

func TestListCodec_NativeToBinary_UInt32(t *testing.T) {
	contents := []uint32{1, 2, 3, 4}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([]uint32{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([]uint32)
	if !ok {
		tests.Failed("Should have received slice of uint []uint")
	}
	tests.Passed("Should have received slice of uint []uint")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}

func TestListCodec_NativeToBinary_UInt64(t *testing.T) {
	contents := []uint64{1, 2, 3, 4}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([]uint64{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([]uint64)
	if !ok {
		tests.Failed("Should have received slice of uint []uint")
	}
	tests.Passed("Should have received slice of uint []uint")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}

func TestListCodec_NativeToBinary_Int(t *testing.T) {
	contents := []int{1, 2, 3, 4}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([]int{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([]int)
	if !ok {
		tests.Failed("Should have received slice of int []int")
	}
	tests.Passed("Should have received slice of int []int")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}

func TestListCodec_NativeToBinary_Int8(t *testing.T) {
	contents := []int8{1, 2, 3, 4}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([]int8{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([]int8)
	if !ok {
		tests.Failed("Should have received slice of int []int")
	}
	tests.Passed("Should have received slice of int []int")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}

func TestListCodec_NativeToBinary_Int16(t *testing.T) {
	contents := []int16{1, 2, 3, 4}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([]int16{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([]int16)
	if !ok {
		tests.Failed("Should have received slice of int []int")
	}
	tests.Passed("Should have received slice of int []int")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}

func TestListCodec_NativeToBinary_Int32(t *testing.T) {
	contents := []int32{1, 2, 3, 4}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([]int32{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([]int32)
	if !ok {
		tests.Failed("Should have received slice of int []int")
	}
	tests.Passed("Should have received slice of int []int")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}

func TestListCodec_NativeToBinary_Int64(t *testing.T) {
	contents := []int64{1, 2, 3, 4}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([]int64{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([]int64)
	if !ok {
		tests.Failed("Should have received slice of int []int")
	}
	tests.Passed("Should have received slice of int []int")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}

func TestListCodec_NativeToBinary_Float32(t *testing.T) {
	contents := []float32{1, 2, 3, 4}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([]float32{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([]float32)
	if !ok {
		tests.Failed("Should have received slice of float []float")
	}
	tests.Passed("Should have received slice of float []float")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}

func TestListCodec_NativeToBinary_Float64(t *testing.T) {
	contents := []float64{1, 2, 3, 4}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([]float64{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([]float64)
	if !ok {
		tests.Failed("Should have received slice of float []float")
	}
	tests.Passed("Should have received slice of float []float")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}

func TestListCodec_NativeToBinary_Bool(t *testing.T) {
	contents := []bool{true, false, true, false}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([]bool{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([]bool)
	if !ok {
		tests.Failed("Should have received slice of bool []bool")
	}
	tests.Passed("Should have received slice of bool []bool")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}

func TestListCodec_NativeToBinary_Text(t *testing.T) {
	contents := []string{"wreckage", "went into downtown", "moppers guild", "God is Love"}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([]string{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([]string)
	if !ok {
		tests.Failed("Should have received slice of string []string")
	}
	tests.Passed("Should have received slice of string []string")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}

func TestListCodec_NativeToBinary_NestedList(t *testing.T) {
	contents := [][]string{
		[]string{"wreckage", "went into downtown"},
		[]string{"moppers guild", "God is Love"},
		[]string{"Is His always Faithful!"},
	}

	var codec codecs.ListCodec
	encoded, err := codec.NativeToBinary(contents, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded value with list codec")
	}
	tests.Passed("Should have successfully encoded value with list codec")

	if jsonEncoded, err := json.Marshal(contents); err == nil {
		tests.Info("JSON Encoded Length: %d", len(jsonEncoded))
		tests.Info("Voxa Encoded Length: %d", len(encoded))
	}

	response, err := codec.BinaryToNative(encoded, reflect.ValueOf([][]string{}))
	if err != nil {
		tests.FailedWithError(err, "Should have successfully decoded value with list codec")
	}
	tests.Passed("Should have successfully decoded value with list codec")

	res, ok := response.([][]string)
	if !ok {
		tests.Failed("Should have received slice of string [][]string")
	}
	tests.Passed("Should have received slice of string [][]string")

	if !reflect.DeepEqual(res, contents) {
		tests.Failed("Should have matching elements between input and output")
	}
	tests.Passed("Should have matching elements between input and output")
}
