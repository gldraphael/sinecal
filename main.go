package main

import (
	"fmt"
	"os"
)

func main() {
	tune, err := os.ReadFile("./assets/tunes/main.tune")
	if err != nil {
		fmt.Println(err)
		return
	}

	Play(ParseTuneFromBytes(tune))
}
