package codecs_test

import (
	"testing"

	"github.com/influx6/faux/tests"
	"github.com/wirekit/voxa/codecs"
)

func TestFloat64_Encoding_Decoding_Unsigned(t *testing.T) {
	val := float64(32.545)
	encoded := codecs.EncodeFloat64(val)
	decoded := codecs.DecodeFloat64(encoded)

	if !float64Equals(val, decoded) {
		tests.Failed("Should have successfully encoded and decoded unsighed float64")
	}
	tests.Passed("Should have successfully encoded and decoded unsighed float64")
}

func TestFloat64_Encoding_Decoding_Signed(t *testing.T) {
	val := float64(-32.545)
	encoded := codecs.EncodeFloat64(val)
	decoded := codecs.DecodeFloat64(encoded)

	if !float64Equals(val, decoded) {
		tests.Failed("Should have successfully encoded and decoded unsighed float64")
	}
	tests.Passed("Should have successfully encoded and decoded unsighed float64")
}

func TestFloat32_Encoding_Decoding_Unsigned(t *testing.T) {
	val := float32(32.5454)
	encoded := codecs.EncodeFloat32(val)
	decoded := codecs.DecodeFloat32(encoded)

	if !float32Equals(val, decoded) {
		tests.Failed("Should have successfully encoded and decoded unsighed float64")
	}
	tests.Passed("Should have successfully encoded and decoded unsighed float64")
}

func TestFloat32_Encoding_Decoding_Signed(t *testing.T) {
	val := float32(-32.5454)
	encoded := codecs.EncodeFloat32(val)
	decoded := codecs.DecodeFloat32(encoded)

	if !float32Equals(val, decoded) {
		tests.Failed("Should have successfully encoded and decoded unsighed float64")
	}
	tests.Passed("Should have successfully encoded and decoded unsighed float64")
}

var EPSILON64 float64 = 0.00000001
var EPSILON32 float32 = 0.00000001

func float32Equals(a, b float32) bool {
	if (a-b) < EPSILON32 && (b-a) < EPSILON32 {
		return true
	}
	return false
}

func float64Equals(a, b float64) bool {
	if (a-b) < EPSILON64 && (b-a) < EPSILON64 {
		return true
	}
	return false
}
