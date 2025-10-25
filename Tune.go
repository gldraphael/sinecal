package main

import (
	"strconv"
	"strings"
)

type Tune []Note

func ParseTuneFromBytes(tune []byte) Tune {
	return ParseTune(string(tune))
}

/*
* Parses a string representation of a tune.
*
* The representation is a sequence of note-duration pairs separated by spaces.
* Each note is represented by a letter (A-G) followed by an optional
* accidental (sharp or flat) and an octave number (0-9).
* Each duration is represented by a number (1, 2, 4, 8, 16, 32 i.e. whole to sixteenth note) followed by a dot (.) for dotted notes.
* The note and duration are separated by a hyphen (-).
* If the duration is not specified, it defaults to 4 (quarter note).
* If a note is not specified but a duration is, it is tied to the previous note.
*
* Example:
* C4 D4-8 E4-8 -4. D4 C4-1
 */
func ParseTune(tune string) Tune {
	if tune == "" {
		return Tune{}
	}

	tokens := strings.Fields(tune)
	var notes []Note
	var lastPitch *Pitch

	for _, token := range tokens {
		parts := strings.Split(token, "-")

		var notePart, durationPart string
		if len(parts) == 1 {
			notePart = parts[0]
			durationPart = "4" // default duration
		} else {
			notePart = parts[0]
			durationPart = parts[1]
		}

		var pitch *Pitch
		if notePart == "" {
			// Tied note - use previous pitch
			pitch = lastPitch
		} else {
			pitch = Parse(notePart)
			lastPitch = pitch
		}

		// Parse duration
		duration := 4.0 // default
		if durationPart != "" {
			isDotted := strings.HasSuffix(durationPart, ".")
			durationStr := strings.TrimSuffix(durationPart, ".")

			if d, err := strconv.Atoi(durationStr); err == nil {
				duration = float64(d)
				if isDotted {
					// Dotted note multiplies duration by 1.5
					// For integer math: duration = duration * 3 / 2
					// But we need to handle this differently since we're working with note values
					// A dotted quarter (4.) has duration value of 2 (half note + quarter)
					switch duration {
					case 1:
						duration = 1.5 // dotted whole becomes whole + half = 1.5
					case 2:
						duration = 1 // dotted half becomes whole + half = 1
					case 4:
						duration = 2 // dotted quarter becomes half + quarter = 2
					case 8:
						duration = 4 // dotted eighth becomes quarter + eighth = 4
					case 16:
						duration = 8 // dotted sixteenth becomes eighth + sixteenth = 8
					case 32:
						duration = 16 // dotted thirty-second becomes sixteenth + thirty-second = 16
					}
				}
			}
		}

		notes = append(notes, Note{
			Pitch:    pitch,
			Duration: duration,
		})
	}

	return notes
}
