package main

import (
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/generators"
)

type Note struct {
	Pitch    *Pitch
	Duration float64
}

func (note Note) Sine() *beep.Streamer {
	tone, _ := generators.SineTone(SampleRate, note.Pitch.Freq())
	if tone == nil {
		tone, _ = generators.SineTone(SampleRate, note.Pitch.Freq())
	}
	return &tone
}

func (note Note) Num() int {
	seconds := time.Duration(note.Duration) * time.Second
	return SampleRate.N(seconds)
}
