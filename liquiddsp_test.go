package liquiddsp

import "testing"

func TestCountBitErrorsArray(t *testing.T) {
	originalMessage := []uint8{0, 1, 2, 3, 4, 5, 6, 7}
	corruptedMessage := []uint8{1, 1, 2, 3, 4, 5, 6, 7}

	bitErrors, err := CountBitErrorsArray(originalMessage, corruptedMessage)
	if err != nil {
		t.Fatal(err)
	}
	if bitErrors != 1 {
		t.Errorf("Got bit error count = %d, wanted 1", bitErrors)
	}
}
