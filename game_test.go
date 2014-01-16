package main

import (
	"fmt"
	"testing"
)

func TestGameUpdateScore(t *testing.T) {
	g := NewGame()

	g.UpdateScore(1, 2)

	if g.P1Score() != 1 {
		t.Error(fmt.Sprintf("P1 score should be 1 but was %d", g.P1Score()))
	}

	if g.P2Score() != 2 {
		t.Error(fmt.Sprintf("P2 score should be 2 but was %d", g.P2Score()))
	}

	if g.CurrentServer() != 1 {
		t.Error(fmt.Sprintf("Current server should be 1 but was %d", g.CurrentServer()))
	}
}

func TestGameCurrentServerAdvances(t *testing.T) {
	g := NewGame()

	g.UpdateScore(1, 0)
	g.UpdateScore(1, 1)

	if g.CurrentServer() != 2 {
		t.Error(fmt.Sprintf("Current server should be 2 but was %d", g.CurrentServer()))
	}
}

func TestGameWinConditions(t *testing.T) {
	g := NewGame()

	g.UpdateScore(11, 5)

	if !g.Over() {
		t.Error("Game should be over but is not")
	}

	g.UpdateScore(14, 13)
	if g.Over() {
		t.Error("Game should not be over but is")
	}

	g.UpdateScore(15, 13)
	if !g.Over() {
		t.Error("Game should be over but is not")
	}
}

func TestGameWinner(t *testing.T) {
	g := NewGame()

	if g.Winner() != 0 {
		t.Error("Game should not have a winner but does")
	}

	g.UpdateScore(11, 0)
	if g.Winner() != 1 {
		t.Error("Player 1 should have won but did not")
	}

	g.UpdateScore(11, 13)
	if g.Winner() != 2 {
		t.Error("Player 2 should have won but did not")
	}
}

func TestGameRestart(t *testing.T) {
	g := NewGame()

	g.UpdateScore(1, 0)
	g.UpdateScore(1, 1)

	g.Restart(1)

	if g.P1Score() != 0 {
		t.Error("P1 score was not reset")
	}

	if g.P2Score() != 0 {
		t.Error("P2 score was not reset")
	}

	if g.CurrentServer() != 1 {
		t.Error("Server was not reset to player 1")
	}

	g.Restart(2)
	if g.CurrentServer() != 2 {
		t.Error("Server was not reset to player 2")
	}
}
