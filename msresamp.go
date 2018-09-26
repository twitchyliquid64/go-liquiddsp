package liquiddsp

// #cgo LDFLAGS: -lliquid -lm -lc
// #include <liquid/liquid.h>
import "C"
import (
	"math"
)

// MsResamp represents a multi-state arbitrary resampler.
type MsResamp struct {
	rs           C.msresamp_crcf
	resampleRate float32
}

// Execute runs resampling on a sample buffer inBuff, returning the output.
func (rs *MsResamp) Execute(inBuff []complex64) []complex64 {
	out := make([]complex64, int(2*len(inBuff)*int(math.Ceil(float64(rs.resampleRate)))))
	var outWritten C.uint
	C.msresamp_crcf_execute(rs.rs, (*C.complexfloat)(&inBuff[0]), C.uint(len(inBuff)), (*C.complexfloat)(&out[0]), (*C.uint)(&outWritten))
	return out[:uint(outWritten)]
}

// NewMultistageResampler creates a new multi-state arbitrary resampler.
func NewMultistageResampler(resampleRate, stopbandSuppression float32) (*MsResamp, error) {
	a := &MsResamp{
		rs:           C.msresamp_crcf_create(C.float(resampleRate), C.float(stopbandSuppression)),
		resampleRate: resampleRate,
	}
	return a, nil
}

// CloseMultistageResampler frees internal memory used for the given MsResamp object.
func CloseMultistageResampler(rs *MsResamp) error {
	C.msresamp_crcf_destroy(rs.rs)
	rs.rs = nil
	return nil
}
