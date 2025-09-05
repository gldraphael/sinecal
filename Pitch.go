package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Pitch struct {
	Name   string
	Octave int
}

func (p Pitch) String() string {
	return fmt.Sprintf("%s%d", p.Name, p.Octave)
}

func (p Pitch) isValid() error {
	noteName := strings.ToUpper(string(p.Name[0]))

	if noteName < "A" || noteName > "G" {
		return errors.New(fmt.Sprintf("The note name %s must be between A and G", noteName))
	}
	if p.Octave < 0 || p.Octave > 8 {
		return errors.New(fmt.Sprint("The octave %n must be between 0 and 8", p.Octave))
	}

	// Special case: only C is allowed in octave 8, and no notes below A in octave 0
	if p.Octave == 8 && noteName != "C" {
		return errors.New(fmt.Sprintf("The octave %d must be between 0 and 8", p.Octave))
	}
	if p.Octave == 0 && noteName > "B" {
		return errors.New(fmt.Sprint("Octave 0 currently only supports A, A#/Bb, B"))
	}

	// TODO: validate sharps & flats...

	return nil
}

func (p Pitch) Freq() float64 {
	// A4 = 440 Hz as reference
	// Calculate semitones from A4
	noteOffsets := map[string]int{
		"C":  -9,
		"C#": -8, "DB": -8,
		"D":  -7,
		"D#": -6, "EB": -6,
		"E":  -5,
		"F":  -4,
		"F#": -3, "GB": -3,
		"G":  -2,
		"G#": -1, "AB": -1,
		"A":  0,
		"A#": 1, "BB": 1,
		"B": 2,
	}

	noteName := strings.ToUpper(p.Name)
	noteOffset, exists := noteOffsets[noteName]
	if !exists {
		return 0 // Invalid note
	}

	// Calculate total semitones from A4
	octaveOffset := (p.Octave - 4) * 12
	totalSemitones := noteOffset + octaveOffset

	// Calculate frequency using equal temperament formula
	// f = 440 * 2^(n/12) where n is semitones from A4
	return 440.0 * float64(int(1)<<(totalSemitones/12)) * pow2(float64(totalSemitones%12)/12.0)
}

func NewPitch(name string, octave int) (*Pitch, error) {
	pitch := &Pitch{
		Name:   name,
		Octave: octave,
	}
	if err := pitch.isValid(); err != nil {
		return nil, err
	}
	return pitch, nil
}

func ParsePitch(pitch string) (*Pitch, error) {
	if len(pitch) < 2 {
		return nil, errors.New("The pitch must be in the format \"C5\"")
	}

	// Find where the octave number starts
	octaveStartIdx := 1
	if len(pitch) > 2 && (pitch[1] == '#' || pitch[1] == 'b') {
		octaveStartIdx = 2
	}

	// Extract name and octave parts
	name := pitch[:octaveStartIdx]
	octaveStr := pitch[octaveStartIdx:]

	// Parse octave
	octave, err := strconv.Atoi(octaveStr)
	if err != nil {
		return nil, err
	}

	p := &Pitch{
		Name:   name,
		Octave: octave,
	}
	if err := p.isValid(); err != nil {
		return nil, err
	}

	return p, nil
}

// This a complete hack, and I'll need to rethink this
func Parse(pitch string) *Pitch {
	p, _ := ParsePitch(pitch)
	if p == nil {
		p, _ = ParsePitch("C4") // this is what they might have meant, obviously!
	}
	return p
}

// Simple approximation of 2^x for fractional powers
func pow2(x float64) float64 {
	// Simple approximation for 2^x where 0 <= x < 1
	// Using Taylor series approximation
	return 1.0 + x*0.693147 + x*x*0.240227 + x*x*x*0.055504
}
