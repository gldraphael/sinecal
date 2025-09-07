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
	tempo := 88 // TODO: do not hardcode this
	// Calculate seconds based on note duration and tempo
	// tempo is in BPM (beats per minute), assuming quarter note = 1 beat
	// Duration: 1 =whole note, 2 = half note, 4 = quarter note, 8 = eighth note
	beatsPerSecond := float64(tempo) / 60.0
	quarterNoteSeconds := 1.0 / beatsPerSecond
	seconds := quarterNoteSeconds * (4.0 / note.Duration)
	return SampleRate.N(time.Duration(seconds * float64(time.Second)))
}
