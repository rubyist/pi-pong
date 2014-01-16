package main

type Game struct {
	P1Score       int
	P2Score       int
	CurrentServer int
	numServes     int
}

func NewGame() *Game {
	return &Game{CurrentServer: 1}
}

func (g *Game) Over() bool {
	return (g.P1Score >= 11 || g.P2Score >= 11) && scoreDiff(g.P1Score, g.P2Score) >= 2
}

func (g *Game) Winner() int {
	if !g.Over() {
		return 0
	}

	if g.P1Score > g.P2Score {
		return 1
	}
	return 2
}

func (g *Game) UpdateScore(p1, p2 int) {
	if (p1 != g.P1Score) || (p2 != g.P2Score) {
		g.P1Score = p1
		g.P2Score = p2
		g.numServes++

		if g.numServes == 2 {
			g.numServes = 0
			if g.CurrentServer == 1 {
				g.CurrentServer = 2
			} else {
				g.CurrentServer = 1
			}
		}
	}
}

func (g *Game) Restart(server int) {
	g.P1Score = 0
	g.P2Score = 0
	g.CurrentServer = server
	g.numServes = 0
}

func scoreDiff(s1, s2 int) int {
	d := s1 - s2
	if d < 0 {
		return -d
	}
	return d
}
