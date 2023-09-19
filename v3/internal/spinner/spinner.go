/*
Copyright © 2023 github.com/alpkeskin
*/
package spinner

import (
	"time"

	"github.com/theckman/yacspin"
)

type Spinner struct {
	spinner *yacspin.Spinner
}

func New(suffix string) *Spinner {
	cfg := yacspin.Config{
		Frequency:         100 * time.Millisecond,
		CharSet:           yacspin.CharSets[59],
		Suffix:            " " + suffix,
		SuffixAutoColon:   true,
		StopCharacter:     "✓",
		StopColors:        []string{"fgGreen"},
		StopFailCharacter: "✗",
		StopFailColors:    []string{"fgRed"},
	}

	spinner, err := yacspin.New(cfg)

	if err != nil {
		panic(err)
	}

	return &Spinner{
		spinner: spinner,
	}
}

func (s *Spinner) Start() {
	s.spinner.Start()
}

func (s *Spinner) Stop() {
	s.spinner.Stop()
}

func (s *Spinner) StopFail() {
	s.spinner.StopFail()
}

func (s *Spinner) SetMessage(message string) {
	s.spinner.Message(message)
}
