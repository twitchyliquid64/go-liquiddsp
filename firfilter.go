package liquiddsp

// #cgo LDFLAGS: -lliquid -lm -lc
// #include <liquid/liquid.h>
import "C"
import (
	"errors"
	"strconv"
)

// FIRFilterPrototype represents an available FIR filter prototype.
type FIRFilterPrototype uint8

// FIR Filter prototypes.
const (
	KaiserFIRFilter FIRFilterPrototype = C.LIQUID_FIRFILT_KAISER
)

// FIR represents different kinds of FIR filters.
type FIR interface {
	Type() FIRFilterPrototype
	SetScale(x float32)
	Push(x complex64)
	Execute() complex64
	FrequencyResponse(evalFreq float32) complex64
}

// CrcfFIR represents a crcf FIR filter.
type CrcfFIR struct {
	filterType FIRFilterPrototype
	filter     C.firfilt_crcf
}

// Type implements FIR.
func (f *CrcfFIR) Type() FIRFilterPrototype {
	return f.filterType
}

// SetScale implements FIR.
func (f *CrcfFIR) SetScale(x float32) {
	C.firfilt_crcf_set_scale(f.filter, C.float(x))
}

// Push implements FIR.
func (f *CrcfFIR) Push(x complex64) {
	C.firfilt_crcf_push(f.filter, C.complexfloat(x))
}

// Execute implements FIR.
func (f *CrcfFIR) Execute() complex64 {
	var out complex64
	C.firfilt_crcf_execute(f.filter, (*C.complexfloat)(&out))
	return out
}

// FrequencyResponse implements FIR.
func (f *CrcfFIR) FrequencyResponse(evalFreq float32) complex64 {
	var out complex64
	C.firfilt_crcf_freqresponse(f.filter, C.float(evalFreq), (*C.complexfloat)(&out))
	return out
}

// NewKaiserFIR initializes a kaiser filter with the given parameters.
//  cutoffFrequency: 0 < cutoffFrequency < 0.5
//  stopbandAttenuation:  _As > 0
//  fractionalSampleOffset: -0.5 < _mu < 0.5
func NewKaiserFIR(filterLength uint32, cutoffFrequency, stopbandAttenuation, fractionalSampleOffset float32) (FIR, error) {
	if filterLength == 0 {
		return nil, errors.New("filter len must be > 0")
	}

	return &CrcfFIR{
		filterType: KaiserFIRFilter,
		filter:     C.firfilt_crcf_create_kaiser(C.uint(filterLength), C.float(cutoffFrequency), C.float(stopbandAttenuation), C.float(fractionalSampleOffset)),
	}, nil
}

// CloseFIR destroys the FIR object.
func CloseFIR(fir FIR) error {
	switch v := fir.(type) {
	case *CrcfFIR:
		C.firfilt_crcf_destroy(v.filter)
		v.filter = nil
		v.filterType = 0
	default:
		return errors.New("unknown FIR: " + strconv.Itoa(int(v.Type())))
	}
	return nil
}
