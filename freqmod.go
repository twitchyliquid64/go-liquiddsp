package liquiddsp

// #cgo LDFLAGS: -lliquid -lm -lc
// #include <liquid/liquid.h>
import "C"

// FreqMod represents an analog FM modulator.
type FreqMod struct {
	fm C.freqmod
}

// Modulate applies the modulator over the individual signal sample.
func (fm *FreqMod) Modulate(signal float32) complex64 {
	var out complex64
	C.freqmod_modulate(fm.fm, C.float(signal), (*C.complexfloat)(&out))
	return out
}

// ModulateBlock applies the modulator over a set of signal samples.
func (fm *FreqMod) ModulateBlock(signal []float32) []complex64 {
	out := make([]complex64, len(signal))
	C.freqmod_modulate_block(fm.fm, (*C.float)(&signal[0]), C.uint(len(signal)), (*C.complexfloat)(&out[0]))
	return out
}

// Reset resets the internal state of the FM modulator.
func (fm *FreqMod) Reset() {
	C.freqmod_reset(fm.fm)
}

// NewFreqMod creates a new FreqMod object.
func NewFreqMod(modulationFactor float32) (*FreqMod, error) {
	fm := C.freqmod_create(C.float(modulationFactor))
	return &FreqMod{
		fm: fm,
	}, nil
}

// CloseFreqMod destroys the internal resources associated with a FreqMod object.
func CloseFreqMod(fm *FreqMod) error {
	C.freqmod_destroy(fm.fm)
	fm.fm = nil
	return nil
}
