package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"

	liquiddsp "github.com/twitchyliquid64/go-liquiddsp"
)

// go run cmd/tonegen/main.go | aplay -f S16_LE -r 44100 --buffer-size=1500

const (
	sampleRate = 44100
)

func dialTone() {
	buff := liquiddsp.MakeTone(sampleRate, sampleRate, 350)
	buff2 := liquiddsp.MakeTone(sampleRate, sampleRate, 450)

	amp := float32(math.Pow(2, 15)) * 0.4

	for i := range buff {
		sample := int16((real(buff[i]) + real(buff2[i])) * amp)
		fmt.Fprintf(os.Stderr, "%5d: %v %5f\n", i, sample, real(buff[i]))
		binary.Write(os.Stdout, binary.LittleEndian, sample)
	}
}

func aussieRingTone() {
	burst1 := liquiddsp.MakeTone(sampleRate*0.4, sampleRate, 425)
	burst2 := liquiddsp.MakeTone(sampleRate*0.4, sampleRate, 450)
	burst3 := liquiddsp.MakeTone(sampleRate*0.4, sampleRate, 400)

	amp := float32(math.Pow(2, 15)) * 0.3

	burst := func() {
		for i := range burst1 {
			sample := int16((real(burst1[i]) + real(burst2[i]) + real(burst3[i])) * amp)
			binary.Write(os.Stdout, binary.LittleEndian, sample)
		}
	}
	pause := func(secs float64) {
		fmt.Fprintf(os.Stderr, "Starting pause\n")
		for i := 0; i < int(sampleRate*secs); i++ {
			os.Stdout.Write([]byte{0, 0})
		}
	}

	for i := 0; i < 4; i++ {
		burst()
		pause(0.2)
		burst()
		pause(2)
	}
}

func main() {
	aussieRingTone()
}
