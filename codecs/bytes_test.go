package codecs_test

import (
	"bytes"
	"testing"

	"github.com/influx6/faux/tests"
	"github.com/wirekit/voxa"
	"github.com/wirekit/voxa/codecs"
)

var (
	bytesTextual     = []byte("we wonder endlessly")
	goodEncodedBytes = append([]byte{byte(voxa.Bytes), 1}, bytesTextual...)
	badEncodedBytes  = append([]byte{byte(voxa.Invalid), 1}, bytesTextual...)
)

func TestBytesCodec_BinaryToNative(t *testing.T) {
	var codec codecs.BytesCodec
	decoded, _, err := codec.BinaryToNative(goodEncodedBytes)
	if err != nil {
		tests.FailedWithError(err, "expected no error with decoding")
	}

	msg, ok := decoded.([]byte)
	if !ok {
		tests.Failed("expected to receive type 'bool'")
	}

	if !bytes.Equal(msg, goodEncodedBytes[2:]) {
		tests.Info("Received: %+q", msg)
		tests.Info("Expected: %+q", goodEncodedBytes[2:])
		tests.Failed("Should have decoded value")
	}
	tests.Passed("Should have decoded value")

	if _, _, err = codec.BinaryToNative(badEncodedBytes); err == nil {
		tests.Failed("expected an error with decoding")
	}
}

func TestBytesCodec_NativeToBinary(t *testing.T) {
	var codec codecs.BytesCodec
	encoded, err := codec.NativeToBinary(bytesTextual, 1, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded boolean value")
	}

	if !bytes.Equal(encoded, goodEncodedBytes) {
		tests.Info("Received: %+q", encoded)
		tests.Info("Expected: %+q", goodEncodedBytes)
		tests.FailedWithError(err, "Should have matched encoded boolean value with expected")
	}
}

func TestBytesCodec_TextualToNative(t *testing.T) {
	var codec codecs.BytesCodec
	decoded, err := codec.TextualToNative(bytesTextual)
	if err != nil {
		tests.FailedWithError(err, "expected no error with decoding")
	}

	msg, ok := decoded.([]byte)
	if !ok {
		tests.Failed("expected to receive type 'bool'")
	}

	if !bytes.Equal(msg, bytesTextual) {
		tests.Info("Received: %+q", msg)
		tests.Info("Expected: %+q", bytesTextual)
		tests.Failed("Should have decoded value")
	}
	tests.Passed("Should have decoded value")
}

func TestBytesCodec_NativeToTextual(t *testing.T) {
	var codec codecs.BytesCodec
	encoded, err := codec.NativeToTextual(bytesTextual, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded boolean value")
	}

	if !bytes.Equal(encoded, bytesTextual) {
		tests.Info("Received: %+q", encoded)
		tests.Info("Expected: %+q", bytesTextual)
		tests.FailedWithError(err, "Should have matched encoded boolean value with expected")
	}
}
