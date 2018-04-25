package codecs_test

import (
	"bytes"
	"testing"

	"github.com/influx6/faux/tests"
	"github.com/wirekit/voxa"
	"github.com/wirekit/voxa/codecs"
)

var (
	intValue       = int32(26)
	intTextual     = []byte("26")
	intBytes       = codecs.EncodeVarInt32(uint32(intValue))
	goodEncodedInt = append([]byte{byte(voxa.Int32), 1}, intBytes...)
	badEncodedInt  = append([]byte{byte(voxa.Invalid), 1}, intBytes...)
)

func TestIntCodec_BinaryToNative(t *testing.T) {
	var codec codecs.IntCodec
	decoded, _, err := codec.BinaryToNative(goodEncodedInt)
	if err != nil {
		tests.FailedWithError(err, "expected no error with decoding")
	}

	value, ok := decoded.(int32)
	if !ok {
		tests.Failed("expected to receive type 'bool'")
	}

	if value != intValue {
		tests.Failed("Should have received expected decoded value")
	}
	tests.Passed("Should have received expected decoded value")

	if _, _, err = codec.BinaryToNative(badEncodedInt); err == nil {
		tests.Failed("expected an error with decoding")
	}
}

func TestIntCodec_NativeToBinary(t *testing.T) {
	var codec codecs.IntCodec
	encoded, err := codec.NativeToBinary(intValue, 1, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded boolean value")
	}

	if !bytes.Equal(encoded, goodEncodedInt) {
		tests.Info("Received: %+q", encoded)
		tests.Info("Expected: %+q", encodedTrue)
		tests.FailedWithError(err, "Should have matched encoded boolean value with expected")
	}
}

func TestIntCodec_TextualToNative(t *testing.T) {
	var codec codecs.IntCodec
	decoded, err := codec.TextualToNative(intTextual)
	if err != nil {
		tests.FailedWithError(err, "expected no error with decoding")
	}

	value, ok := decoded.(int64)
	if !ok {
		tests.Failed("expected to receive type 'bool'")
	}

	if int32(value) != intValue {
		tests.Failed("Should have received expected decoded value")
	}
	tests.Passed("Should have received expected decoded value")
}

func TestIntCodec_NativeToTextual(t *testing.T) {
	var codec codecs.IntCodec
	encoded, err := codec.NativeToTextual(intValue, []byte{})
	if err != nil {
		tests.FailedWithError(err, "Should have successfully encoded boolean value")
	}

	if !bytes.Equal(encoded, intTextual) {
		tests.Info("Received: %+q", encoded)
		tests.Info("Expected: %+q", intTextual)
		tests.FailedWithError(err, "Should have matched encoded boolean value with expected")
	}
}

func must(err error) bool {
	if err != nil {
		panic(err)
	}
	return false
}
