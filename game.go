package main

type Game struct {
	p1Score       int
	p2Score       int
	currentServer int
	numServes     int
}

func NewGame() *Game {
	return &Game{currentServer: 1}
}

func (g *Game) P1Score() int {
	return g.p1Score
}

func (g *Game) P2Score() int {
	return g.p2Score
}

func (g *Game) CurrentServer() int {
	return g.currentServer
}

func (g *Game) Over() bool {
	return (g.P1Score() >= 11 || g.P2Score() >= 11) && scoreDiff(g.P1Score(), g.P2Score()) >= 2
}

func (g *Game) Winner() int {
	if !g.Over() {
		return 0
	}

	if g.P1Score() > g.P2Score() {
		return 1
	}
	return 2
}

func (g *Game) UpdateScore(p1, p2 int) {
	if (p1 != g.P1Score()) || (p2 != g.P2Score()) {
		g.p1Score = p1
		g.p2Score = p2
		g.numServes++

		if g.numServes == 2 {
			g.numServes = 0
			if g.CurrentServer() == 1 {
				g.currentServer = 2
			} else {
				g.currentServer = 1
			}
		}
	}
}

func (g *Game) Restart(server int) {
	g.p1Score = 0
	g.p2Score = 0
	g.currentServer = server
	g.numServes = 0
}

func scoreDiff(s1, s2 int) int {
	d := s1 - s2
	if d < 0 {
		return -d
	}
	return d
}
