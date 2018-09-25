package liquiddsp

import "testing"

func TestBasicAGC(t *testing.T) {
	agc, err := NewAGC()
	if err != nil {
		t.Fatal(err)
	}
	defer CloseAGC(agc)

	if agc.GetGain() != 1 {
		t.Errorf("Expected gain = %v, got %v", 1, agc.GetGain())
	}

	if agc.GetSignalLevel() != 1 {
		t.Errorf("Expected sig = %v, got %v", 1, agc.GetSignalLevel())
	}

	x := agc.Execute(complex(1, 0))
	if x != complex(1, 0) {
		t.Errorf("Expected 1 + 0i, got %v", x)
	}
}
