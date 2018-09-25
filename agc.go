package liquiddsp

// #cgo LDFLAGS: -lliquid -lm -lc
// #include <liquid/liquid.h>
import "C"

// AGC represents automatic gain control.
type AGC struct {
	agc C.agc_crcf
}

// Reset clears the internal state of the AGC, unlocking it and clearing the signal level estimate.
func (a *AGC) Reset() {
	C.agc_crcf_reset(a.agc)
}

// // SetGainLimits sets the minimum and maximum gain values, beyond the window specified by min/max the agc adjusts the gain.
// func (a *AGC) SetGainLimits(min, max float32) {
// 	C.agc_crcf_set_gain_limits(a.agc, min, max)
// }

//TODO: Implement squelch, lock functions.

// Execute provides a sample to the AGC, returns an updated sample, and updates the internal state of the AGC.
func (a *AGC) Execute(input complex64) complex64 {
	var out complex64
	C.agc_crcf_execute(a.agc, C.complexfloat(input), (*C.complexfloat)(&out))
	return out
}

// GetSignalLevel returns a linear estimate of the input signal's energy level.
func (a *AGC) GetSignalLevel() float32 {
	return float32(C.agc_crcf_get_signal_level(a.agc))
}

// GetRSSI returns an estimate of the input signal's energy level in dB.
func (a *AGC) GetRSSI() float32 {
	return float32(C.agc_crcf_get_rssi(a.agc))
}

// GetGain returns the AGCs's internal gain.
func (a *AGC) GetGain() float32 {
	return float32(C.agc_crcf_get_gain(a.agc))
}

// NewAGC creates a new complex AGC.
func NewAGC() (*AGC, error) {
	agc := C.agc_crcf_create()
	return &AGC{
		agc: agc,
	}, nil
}

// CloseAGC destroys the resources/state associated with the AGC.
func CloseAGC(agc *AGC) error {
	C.agc_crcf_destroy(agc.agc)
	agc.agc = nil
	return nil
}
