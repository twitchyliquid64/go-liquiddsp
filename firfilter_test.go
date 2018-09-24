package liquiddsp

import (
	"math"
	"math/cmplx"
	"testing"
)

const cPi = 3.14159265358979323846264338327950288

// Port of: https://github.com/jgaeddert/liquid-dsp/blob/9658d811f9194229304fec2d117f49c59b49a616/examples/firfilt_crcf_example.c
func TestKaiserFilter(t *testing.T) {
	var cutoff float32 = 0.1
	numSamples := 240

	fltr, err := NewKaiserFIR(65, cutoff, 60.0, 0.0)
	if err != nil {
		t.Fatal(err)
	}
	defer CloseFIR(fltr)
	fltr.SetScale(2.0 * cutoff)

	x := make([]complex64, numSamples)
	y := make([]complex64, numSamples)

	wlen := math.Round(0.75 * float64(numSamples))
	for i := 0; i < numSamples; i++ {
		x[i] = complex64(0.7*cmplx.Exp(complex128(complex(0, 2*cPi*0.057*float32(i)))) + 0.3*cmplx.Exp(complex128(complex(0, 2*cPi*0.357*float32(i)))))

		if float64(i) < wlen {
			x[i] *= complex(Hamming(uint(i), uint(wlen)), 0)
		} else {
			x[i] *= complex(0, 0)
		}

		fltr.Push(x[i])
		y[i] = fltr.Execute()
	}

	//TODO: Check values.
}
