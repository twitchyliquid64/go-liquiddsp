package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"

	liquiddsp "github.com/twitchyliquid64/go-liquiddsp"
)

// go run cmd/tonegen/main.go | aplay -f S16_LE -r 44100

const (
	sampleRate = 44100
)

func main() {
	buff := liquiddsp.MakeTone(sampleRate, sampleRate, 440)

	amp := float32(math.Pow(2, 15)) * 0.8

	for i := range buff {
		sample := int16(real(buff[i]) * amp)
		fmt.Fprintf(os.Stderr, "%5d: %v %5f\n", i, sample, real(buff[i]))
		binary.Write(os.Stdout, binary.LittleEndian, sample)
	}
}
