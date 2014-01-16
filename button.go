package main

import (
	"github.com/mrmorphic/hwio"
)

type ButtonCallback func()
type ButtonChangeCallback func(state int)

type Button struct {
	pin     hwio.Pin
	Rising  ButtonCallback
	Falling ButtonCallback
	Change  ButtonChangeCallback
}

func NewButton(pin string) (*Button, error) {
	buttonPin, err := hwio.GetPin(pin)
	if err != nil {
		return nil, err
	}

	err = hwio.PinMode(buttonPin, hwio.INPUT)
	if err != nil {
		return nil, err
	}

	return &Button{pin: buttonPin}, nil
}

func (b *Button) Poll() {
	go func() {
		lastState, _ := hwio.DigitalRead(b.pin)
		for {
			currentState, _ := hwio.DigitalRead(b.pin)
			if currentState != lastState {
				lastState = currentState
				if currentState == 1 {
					if b.Rising != nil {
						b.Rising()
					}
				} else {
					if b.Falling != nil {
						b.Falling()
					}
				}

				if b.Change != nil {
					b.Change(currentState)
				}
			}
			hwio.Delay(50)
		}
	}()
}
