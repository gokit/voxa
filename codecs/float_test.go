package codecs_test

import (
	"bytes"
	"testing"

	"github.com/influx6/faux/tests"
	"github.com/wirekit/voxa"
	"github.com/wirekit/voxa/codecs"
)

var (
	floatValue       float32 = 2.146
	floatTextual             = []byte("2.146")
	floatBytes               = codecs.EncodeVarInt32(codecs.EncodeFloat32(2.146))
	goodEncodedFloat         = append([]byte{byte(voxa.Float32), 1}, floatBytes...)
	badEncodedFloat          = append([]byte{byte(voxa.Invalid), 1}, floatTextual...)
)

func TestFloatCodec_BinaryToNative(t *testing.T) {
	var codec codecs.FloatCodec
	decoded, _, err := codec.BinaryToNative(goodEncodedFloat)
	if err != nil {
		tests.FailedWithError(err, "expected no error with decoding")
	}

	value, ok := decoded.(float32)
	if !ok {
		tests.Failed("expected to receive type 'bool'")
	}

	if !float32Equals(floatValue, value) {
		tests.Failed("Should have received decoded value")
	}
	tests.Passed("Should have received expected decoded value")

	if _, _, err = codec.BinaryToNative(badEncodedFloat); err == nil {
		tests.Failed("expected an error with decoding")
	}
}

func TestFloatCodec_NativeToBinary(t *testing.T) {
	var codec codecs.FloatCodec
	encoded, err := codec.NativeToBinary(floatValue, 1, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded boolean value")
	}

	if !bytes.Equal(encoded, goodEncodedFloat) {
		tests.Info("Received: %+q", encoded)
		tests.Info("Expected: %+q", encodedTrue)
		tests.FailedWithError(err, "Should have matched encoded boolean value with expected")
	}
}

func TestFloatCodec_TextualToNative(t *testing.T) {
	var codec codecs.FloatCodec
	decoded, err := codec.TextualToNative(floatTextual)
	if err != nil {
		tests.FailedWithError(err, "expected no error with decoding")
	}

	value, ok := decoded.(float64)
	if !ok {
		tests.Failed("expected to receive type 'float64/float32'")
	}

	if !float32Equals(float32(value), floatValue) {
		tests.Info("Received: %#v", value)
		tests.Info("Expected: %#v", floatValue)
		tests.Failed("Should have received expected decoded value")
	}
	tests.Passed("Should have received expected decoded value")
}

func TestFloatCodec_NativeToTextual(t *testing.T) {
	var codec codecs.FloatCodec
	encoded, err := codec.NativeToTextual(floatValue, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded boolean value")
	}

	expected := []byte("2.1459999084")
	if !bytes.Equal(encoded, expected) {
		tests.Info("Received: %+q", encoded)
		tests.Info("Expected: %+q", expected)
		tests.FailedWithError(err, "Should have matched encoded boolean value with expected")
	}
}
