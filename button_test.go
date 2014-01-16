package main

import (
	"fmt"
	"github.com/mrmorphic/hwio"
	"testing"
)

func TestButtonCreation(t *testing.T) {
	driver := new(hwio.TestDriver)
	hwio.SetDriver(driver)

	_, err := NewButton("P1")

	if err != nil {
		t.Error(fmt.Sprintf("Button creation failed: %s", err))
	}

	_, err = NewButton("P99")

	if err == nil {
		t.Error("Create button should have failed but didn't")
	}
}

func TestButtonRising(t *testing.T) {
	driver := new(hwio.TestDriver)
	hwio.SetDriver(driver)

	button, err := NewButton("P1")
	if err != nil {
		t.Error(fmt.Sprintf("Button creation failed: %s", err))
	}

	risingCalled := false
	button.Rising = func() {
		risingCalled = true
	}

	button.Poll()

	// Write to the pin I guess?

	if risingCalled != true {
		t.Error("Expected Rising to be called but it wasn't")
	}
}

// hwio does not export its most useful mocks. Might need to fork it.
