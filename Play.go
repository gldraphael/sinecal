package main

import (
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/speaker"
)

func Play(tune Tune) {
	speaker.Init(SampleRate, 4800)

	notes := tune

	ch := make(chan struct{})
	sounds := make([]beep.Streamer, len(notes))
	for i, n := range notes {
		sounds[i] = beep.Take(n.Num(), *n.Sine())
	}
	sounds = append(sounds, beep.Callback(func() {
		ch <- struct{}{}
	}))

	speaker.Play(beep.Seq(sounds...))
	<-ch
}
