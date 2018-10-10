package liquiddsp

import "math"

// MakeTone produces complex samples that represent a sinosoid at a specific frequency & sample rate.
func MakeTone(numSamples int, sampleRate, frequency float32) []complex64 {
	radialFreq := 2 * math.Pi * frequency / sampleRate
	out := make([]complex64, numSamples)

	var phase float64
	for i := 0; i < numSamples; i++ {
		out[i] = complex64(complex(math.Cos(phase), math.Sin(phase)))
		phase += float64(radialFreq)

		if phase > math.Pi {
			phase -= 2 * math.Pi
		}
		if phase < -math.Pi {
			phase += 2 * math.Pi
		}
	}
	return out
}
