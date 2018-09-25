package liquiddsp

// #cgo LDFLAGS: -lliquid -lm -lc
// #include <liquid/liquid.h>
import "C"

// MultiSource represents a configurable, generated signal source, based on the combination of a number of primitive elements.
type MultiSource struct {
	ms       C.msourcecf
	features map[int]*MultiSourceFeature
}

// MultiSourceFeatureType describes the kind of signal a MultiSourceFeature represents.
type MultiSourceFeatureType uint8

// known types.
const (
	MSToneFeature MultiSourceFeatureType = iota
	MSNoiseFeature
)

// MultiSourceFeature represents a primitive element of a MultiSource.
type MultiSourceFeature struct {
	kind   MultiSourceFeatureType
	id     int
	parent *MultiSource
}

// NewMultiSource creates a new, unconfigured MultiSource.
func NewMultiSource() (*MultiSource, error) {
	ms := C.msourcecf_create()
	return &MultiSource{
		ms:       ms,
		features: map[int]*MultiSourceFeature{},
	}, nil
}

// CloseMultiSource destroys the internal state and buffer resources associated with a MultiSource.
func CloseMultiSource(ms *MultiSource) error {
	C.msourcecf_destroy(ms.ms)
	ms.ms = nil
	return nil
}

// ToneSignal creates a new tone-based signal.
func (ms *MultiSource) ToneSignal() (*MultiSourceFeature, error) {
	s := int(C.msourcecf_add_tone(ms.ms))
	msf := &MultiSourceFeature{
		kind:   MSToneFeature,
		id:     s,
		parent: ms,
	}
	ms.features[s] = msf
	return msf, nil
}

// NoiseSignal creates a new noise signal at a given bandwidth.
func (ms *MultiSource) NoiseSignal(bandwidth float32) (*MultiSourceFeature, error) {
	s := int(C.msourcecf_add_noise(ms.ms, C.float(bandwidth)))
	msf := &MultiSourceFeature{
		kind:   MSNoiseFeature,
		id:     s,
		parent: ms,
	}
	ms.features[s] = msf
	return msf, nil
}

// GetSamples generates and returns the requested number of samples.
func (ms *MultiSource) GetSamples(buffSize int) []complex64 {
	out := make([]complex64, buffSize)
	C.msourcecf_write_samples(ms.ms, (*C.complexfloat)(&out[0]), C.uint(buffSize))
	return out
}

// SetGain sets the signal gain in dB.
func (msf *MultiSourceFeature) SetGain(gain float64) error {
	C.msourcecf_set_gain(msf.parent.ms, C.int(msf.id), C.float(gain))
	return nil
}

// SetFrequency sets the signal's angular frequency relative to the sample rate. The frequency must be within bounds of -Pi to +Pi.
func (msf *MultiSourceFeature) SetFrequency(dphi float64) error {
	C.msourcecf_set_frequency(msf.parent.ms, C.int(msf.id), C.float(dphi))
	return nil
}

// Enable turns on the signal.
func (msf *MultiSourceFeature) Enable() {
	C.msourcecf_enable(msf.parent.ms, C.int(msf.id))
}

// Disable turns off the signal.
func (msf *MultiSourceFeature) Disable() {
	C.msourcecf_disable(msf.parent.ms, C.int(msf.id))
}
