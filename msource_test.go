package liquiddsp

import (
	"math"
	"testing"
)

func TestBasicMsource(t *testing.T) {
	ms, err := NewMultiSource()
	if err != nil {
		t.Fatal(err)
	}
	defer CloseMultiSource(ms)

	tone, err := ms.ToneSignal()
	if err != nil {
		t.Fatal(err)
	}
	if tone.kind != MSToneFeature {
		t.Errorf("Got tone.kind = %v, expected %v", tone.kind, MSToneFeature)
	}
	tone.SetFrequency(-0.4 * 2 * math.Pi)
	tone.SetGain(-40)
	tone.Enable()

	buff := ms.GetSamples(12)
	if len(buff) != 12 {
		t.Fatalf("Expected len(buff) = 12, got %d", len(buff))
	}
	if buff[0] != complex(0.01, 0) {
		t.Errorf("Expected buff[0] = 0.01+0i, got %v", buff[0])
	}
}
