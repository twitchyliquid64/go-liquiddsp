package liquiddsp

import (
	"reflect"
	"testing"
)

func TestBasicFEC(t *testing.T) {
	var tcs = []struct {
		scheme      FECScheme
		codec, name string
		encMsgLen   uint
	}{
		{
			scheme:    HammingFEC74,
			codec:     "h74",
			name:      "Hamming(7,4)",
			encMsgLen: 14,
		},
		{
			scheme:    HammingFEC84,
			codec:     "h84",
			name:      "Hamming(8,4)",
			encMsgLen: 16,
		},
		{
			scheme:    HammingFEC128,
			codec:     "h128",
			name:      "Hamming(12,8)",
			encMsgLen: 12,
		},
		{
			scheme:    Rep3FEC,
			codec:     "rep3",
			name:      "repeat(3)",
			encMsgLen: 24,
		},
		{
			scheme:    Rep5FEC,
			codec:     "rep5",
			name:      "repeat(5)",
			encMsgLen: 40,
		},
		{
			scheme:    NoneFEC,
			codec:     "none",
			name:      "none",
			encMsgLen: 8,
		},
		// error: fec_create(), Reed-Solomon codes unavailable (install libfec)
		//{
		//	scheme: ReedSolomonFEC,
		//	codec:  "h74",
		//},
	}

	for i, tc := range tcs {
		fec, err := CreateFEC(tc.scheme)
		if err != nil {
			t.Errorf("Failed to create FEC in testcase %d", i)
			continue
		}
		defer CloseFEC(fec)

		if fec.Codec() != tc.codec {
			t.Errorf("Codec() = %q, want %q", fec.Codec(), tc.codec)
		}
		if fec.Name() != tc.name {
			t.Errorf("Name() = %q, want %q", fec.Name(), tc.name)
		}
		if fec.EncodedMsgLen(8) != tc.encMsgLen {
			t.Errorf("EncodedMsgLen(8) = %d, want %d", fec.EncodedMsgLen(8), tc.encMsgLen)
		}
	}
}

func TestEncodeDecodeFEC(t *testing.T) {
	var tcs = []struct {
		scheme  FECScheme
		in, out []uint8
	}{
		{
			scheme: HammingFEC74,
			in:     []uint8{4, 8, 12, 55},
			out:    []uint8{1, 48, 7, 0, 15, 33, 143},
		},
		{
			scheme: Rep5FEC,
			in:     []uint8{55},
			out:    []uint8{55, 55, 55, 55, 55},
		},
	}

	for i, tc := range tcs {
		fec, err := CreateFEC(tc.scheme)
		if err != nil {
			t.Errorf("Failed to create FEC in testcase %d", i)
			continue
		}
		defer CloseFEC(fec)

		enc := fec.Encode(tc.in)
		if !reflect.DeepEqual(enc, tc.out) {
			t.Errorf("%s.Encode(%v) = %v, wanted %v", fec.Codec(), tc.in, enc, tc.out)
		}

		dec := fec.Decode(enc, len(tc.in))
		if !reflect.DeepEqual(dec, tc.in) {
			t.Errorf("%s.Decode(%v) = %v, wanted %v", fec.Codec(), enc, dec, tc.in)
		}
	}
}
