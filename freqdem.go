package liquiddsp

// #cgo LDFLAGS: -lliquid -lm -lc
// #include <liquid/liquid.h>
import "C"

// FMDemod represents an FM demodulator.
type FMDemod struct {
	fm C.freqdem
}

// DemodulateBlock applies the demodulator over the individual complex sample.
func (fm *FMDemod) DemodulateBlock(signal []complex64) []float32 {
	out := make([]float32, len(signal))
	C.freqdem_demodulate_block(fm.fm, (*C.complexfloat)(&signal[0]), C.uint(len(signal)), (*C.float)(&out[0]))
	return out
}

// Demodulate applies the demodulator over the individual complex sample.
func (fm *FMDemod) Demodulate(signal complex64) float32 {
	var out float32
	C.freqdem_demodulate(fm.fm, C.complexfloat(signal), (*C.float)(&out))
	return out
}

// Reset resets the internal state of the FM demodulator.
func (fm *FMDemod) Reset() {
	C.freqdem_reset(fm.fm)
}

// NewFMDemod creates a new FR demodulator with the specified modulation factor.
func NewFMDemod(modulationFactor float32) (*FMDemod, error) {
	fm := C.freqdem_create(C.float(modulationFactor))
	return &FMDemod{
		fm: fm,
	}, nil
}

// CloseFMDemod destroys the internal resources associated with a FMDemod object.
func CloseFMDemod(fm *FMDemod) error {
	C.freqdem_destroy(fm.fm)
	fm.fm = nil
	return nil
}
