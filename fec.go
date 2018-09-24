package liquiddsp

// #cgo LDFLAGS: -lliquid -lm -lc
// #include <liquid/liquid.h>
import "C"
import (
	"errors"
)

// FECScheme represents an available FEC codec.
type FECScheme uint8

// FEC schemes.
const (
	UnknownFEC    FECScheme = C.LIQUID_FEC_UNKNOWN
	NoneFEC       FECScheme = C.LIQUID_FEC_NONE
	HammingFEC74  FECScheme = C.LIQUID_FEC_HAMMING74
	HammingFEC84  FECScheme = C.LIQUID_FEC_HAMMING84
	HammingFEC128 FECScheme = C.LIQUID_FEC_HAMMING128
	Rep3FEC       FECScheme = C.LIQUID_FEC_REP3
	Rep5FEC       FECScheme = C.LIQUID_FEC_REP5

	// NOTE: Must have libfec installed.
	ReedSolomonFEC FECScheme = C.LIQUID_FEC_RS_M8
)

// FEC represents a FEC encoder/decoder.
type FEC struct {
	scheme FECScheme
	fec    C.fec
}

// Codec returns a short string representation of the FEC scheme.
func (f *FEC) Codec() string {
	return C.GoString(C.fec_scheme_str[f.scheme][0])
}

// Name returns the long-form string representation of the FEC scheme.
func (f *FEC) Name() string {
	return C.GoString(C.fec_scheme_str[f.scheme][1])
}

// EncodedMsgLen indicates the size of a message of length msgLen after it has been encoded.
func (f *FEC) EncodedMsgLen(msgLen uint) uint {
	return uint(C.fec_get_enc_msg_length(C.fec_scheme(f.scheme), C.uint(msgLen)))
}

// Encode runs the FEC over msg, returning an encoded message.
func (f *FEC) Encode(msg []uint8) []uint8 {
	out := make([]uint8, f.EncodedMsgLen(uint(len(msg))))
	C.fec_encode(f.fec, C.uint(len(msg)), (*C.uchar)(&msg[0]), (*C.uchar)(&out[0]))
	return out
}

// Decode runs the decodes the message into a buffer of size outputLen, before returning that buffer.
func (f *FEC) Decode(encoded []uint8, outputLen int) []uint8 {
	out := make([]uint8, outputLen)
	C.fec_decode(f.fec, C.uint(outputLen), (*C.uchar)(&encoded[0]), (*C.uchar)(&out[0]))
	return out
}

// CreateFEC initializes a FEC for use.
func CreateFEC(scheme FECScheme) (*FEC, error) {
	var fec C.fec

	switch scheme {
	case NoneFEC:
		fallthrough
	case Rep3FEC:
		fallthrough
	case Rep5FEC:
		fallthrough
	case ReedSolomonFEC:
		fallthrough
	case HammingFEC74:
		fallthrough
	case HammingFEC84:
		fallthrough
	case HammingFEC128:
		fec = C.fec_create(C.fec_scheme(scheme), nil)
	default:
		return nil, errors.New("unrecognised FEC scheme")
	}

	if fec == nil {
		return nil, errors.New("internal error: fec is null")
	}

	return &FEC{fec: fec, scheme: scheme}, nil
}

// CloseFEC destroys a FEC object.
func CloseFEC(fec *FEC) error {
	C.fec_destroy(fec.fec)
	fec.fec = nil
	fec.scheme = UnknownFEC
	return nil
}
