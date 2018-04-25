package codecs_test

import (
	"bytes"
	"testing"

	"github.com/influx6/faux/tests"
	"github.com/wirekit/voxa"
	"github.com/wirekit/voxa/codecs"
)

var (
	trueTextual    = []byte("true")
	falseTextual   = []byte("false")
	encodedTrue    = []byte{byte(voxa.Boolean), 1, 1}
	encodedFalse   = []byte{byte(voxa.Boolean), 1, 0}
	badEncodedTrue = []byte{byte(voxa.Invalid), 1, 1}
)

func TestBoolCodec_BinaryToNative(t *testing.T) {
	var codec codecs.BooleanCodec
	t.Logf("Should be able to decode True value encoding")
	{
		decoded, _, err := codec.BinaryToNative(encodedTrue)
		if err != nil {
			tests.FailedWithError(err, "expected no error with decoding")
		}

		boolValue, ok := decoded.(bool)
		if !ok {
			tests.Failed("expected to receive type 'bool'")
		}

		if !boolValue {
			tests.Failed("Should have received expected decoded value")
		}
		tests.Passed("Should have received expected decoded value")

		if _, _, err = codec.BinaryToNative(badEncodedTrue); err == nil {
			tests.Failed("expected an error with decoding")
		}
	}

	t.Logf("Should be able to decode False value encoding")
	{
		decoded, _, err := codec.BinaryToNative(encodedFalse)
		if err != nil {
			tests.FailedWithError(err, "expected no error with decoding")
		}

		boolValue, ok := decoded.(bool)
		if !ok {
			tests.Failed("expected to receive type 'bool'")
		}

		if boolValue {
			tests.Failed("Should have received false as decoded value")
		}
		tests.Passed("Should have received false as decoded value")
	}
}

func TestBooleanCodec_NativeToBinary(t *testing.T) {
	var codec codecs.BooleanCodec
	t.Log("Should be able to encode boolean value 'true'")
	{
		encoded, err := codec.NativeToBinary(true, 1, []byte{})
		if err != nil {
			tests.FailedWithError(err, "Should have successfully encoded boolean value")
		}

		if !bytes.Equal(encoded, encodedTrue) {
			tests.Info("Received: %+q", encoded)
			tests.Info("Expected: %+q", encodedTrue)
			tests.FailedWithError(err, "Should have matched encoded boolean value with expected")
		}
	}

	t.Log("Should be able to encode boolean value 'false'")
	{
		encoded, err := codec.NativeToBinary(false, 1, []byte{})
		if err != nil {
			tests.FailedWithError(err, "Should have successfully encoded boolean value")
		}

		if !bytes.Equal(encoded, encodedFalse) {
			tests.Info("Received: %+q", encoded)
			tests.Info("Expected: %+q", encodedFalse)
			tests.FailedWithError(err, "Should have matched encoded boolean value with expected")
		}
	}
}

func TestBoolCodec_TextualToNative(t *testing.T) {
	var codec codecs.BooleanCodec
	t.Logf("Should be able to decode True value encoding")
	{
		decoded, err := codec.TextualToNative(trueTextual)
		if err != nil {
			tests.FailedWithError(err, "expected no error with decoding")
		}

		boolValue, ok := decoded.(bool)
		if !ok {
			tests.Failed("expected to receive type 'bool'")
		}

		if !boolValue {
			tests.Failed("Should have received expected decoded value")
		}
		tests.Passed("Should have received expected decoded value")
	}

	t.Logf("Should be able to decode False value encoding")
	{
		decoded, err := codec.TextualToNative(falseTextual)
		if err != nil {
			tests.FailedWithError(err, "expected no error with decoding")
		}

		boolValue, ok := decoded.(bool)
		if !ok {
			tests.Failed("expected to receive type 'bool'")
		}

		if boolValue {
			tests.Failed("Should have received false as decoded value")
		}
		tests.Passed("Should have received false as decoded value")
	}
}

func TestBooleanCodec_NativeToTextual(t *testing.T) {
	var codec codecs.BooleanCodec
	t.Log("Should be able to encode boolean value 'true'")
	{
		encoded, err := codec.NativeToTextual(true, []byte{})
		if err != nil {
			tests.FailedWithError(err, "Should have successfully encoded boolean value")
		}

		if !bytes.Equal(encoded, trueTextual) {
			tests.Info("Received: %+q", encoded)
			tests.Info("Expected: %+q", trueTextual)
			tests.FailedWithError(err, "Should have matched encoded boolean value with expected")
		}
	}

	t.Log("Should be able to encode boolean value 'false'")
	{
		encoded, err := codec.NativeToTextual(false, []byte{})
		if err != nil {
			tests.FailedWithError(err, "Should have successfully encoded boolean value")
		}

		if !bytes.Equal(encoded, falseTextual) {
			tests.Info("Received: %+q", encoded)
			tests.Info("Expected: %+q", falseTextual)
			tests.FailedWithError(err, "Should have matched encoded boolean value with expected")
		}
	}
}
