package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"

	liquiddsp "github.com/twitchyliquid64/go-liquiddsp"
)

func main() {
	buff := liquiddsp.MakeTone(8000, 8000, 440)

	amp := float32(math.Pow(2, 15)) * 0.7

	for i := range buff {
		sample := int16(real(buff[i]) * amp)
		fmt.Fprintf(os.Stderr, "%5d: %v %5f\n", i, sample, real(buff[i]))
		binary.Write(os.Stdout, binary.LittleEndian, sample)
	}
}
