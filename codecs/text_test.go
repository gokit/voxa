package codecs_test

import (
	"bytes"
	"testing"

	"github.com/influx6/faux/tests"
	"github.com/wirekit/voxa"
	"github.com/wirekit/voxa/codecs"
)

var (
	textValue       = "we going to reck some stages"
	textTextual     = []byte("\"" + textValue + "\"")
	goodEncodedText = append([]byte{byte(voxa.Text), 1}, textValue...)
	badEncodedText  = append([]byte{byte(voxa.Invalid), 1}, textValue...)
)

func TestTextCodec_BinaryToNative(t *testing.T) {
	var codec codecs.TextCodec
	decoded, _, err := codec.BinaryToNative(goodEncodedText)
	if err != nil {
		tests.FailedWithError(err, "expected no error with decoding")
	}

	value, ok := decoded.(string)
	if !ok {
		tests.Info("Received: %+q", decoded)
		tests.Failed("expected to receive type 'bool'")
	}

	if value != textValue {
		tests.Info("Received: %+q", value)
		tests.Info("Expected: %+q", textValue)
		tests.Failed("Should have received expected decoded value")
	}
	tests.Passed("Should have received expected decoded value")

	if _, _, err = codec.BinaryToNative(badEncodedText); err == nil {
		tests.Failed("expected an error with decoding")
	}
}

func TestTextCodec_NativeToBinary(t *testing.T) {
	var codec codecs.TextCodec
	encoded, err := codec.NativeToBinary(textValue, 1, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded boolean value")
	}

	if !bytes.Equal(encoded, goodEncodedText) {
		tests.Info("Received: %+q", encoded)
		tests.Info("Expected: %+q", goodEncodedText)
		tests.FailedWithError(err, "Should have matched encoded boolean value with expected")
	}
}

func TestTextCodec_TextualToNative(t *testing.T) {
	var codec codecs.TextCodec
	decoded, err := codec.TextualToNative(textTextual)
	if err != nil {
		tests.FailedWithError(err, "expected no error with decoding")
	}

	value, ok := decoded.(string)
	if !ok {
		tests.Failed("expected to receive type 'bool'")
	}

	if value != string(textValue) {
		tests.Info("Received: %+q", value)
		tests.Info("Expected: %+q", textValue)
		tests.Failed("Should have received expected decoded value")
	}
	tests.Passed("Should have received expected decoded value")
}

func TestTextCodec_NativeToTextual(t *testing.T) {
	var codec codecs.TextCodec
	encoded, err := codec.NativeToTextual(textValue, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded boolean value")
	}

	if !bytes.Equal(encoded, textTextual) {
		tests.Info("Received: %+q", encoded)
		tests.Info("Expected: %+q", textTextual)
		tests.FailedWithError(err, "Should have matched encoded boolean value with expected")
	}
}
