package liquiddsp

// #cgo LDFLAGS: -lliquid -lm -lc
// #include <liquid/liquid.h>
import "C"
import (
	"errors"
)

// CountBitErrorsArray implements liquid-dsp's count_bit_errors_array() function.
func CountBitErrorsArray(original, decoded []uint8) (uint, error) {
	if len(original) != len(decoded) {
		return 0, errors.New("original and decoded messages must be of the same length")
	}
	if len(original) == 0 {
		return 0, errors.New("message buffer cannot be empty")
	}

	return uint(C.count_bit_errors_array((*C.uchar)(&original[0]), (*C.uchar)(&decoded[0]), C.uint(len(decoded)))), nil
}

// Hamming computes the nth of N indices of the hamming window.
func Hamming(nth, n uint) float32 {
	return float32(C.hamming(C.uint(nth), C.uint(n)))
}
