package game

import (
	"fmt"
	"image/color"
	"spaceship_go/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	player           *Player
	lasers           []*Laser
	meteros          []*Meteor
	meteorSpawnTimer *Timer
	score            int
}

func NewGame() *Game {
	g := &Game{
		meteorSpawnTimer: NewTimer(24),
	}
	player := NewPlayer(g)
	g.player = player

	return g
}

// Rodar em 60 FPS.
// Responsavel por atualizar a logica do jogo.
// 60 X por segundo.
// 1 x rodando a cada 1/60 segundos.
func (g *Game) Update() error {
	g.player.Update()

	for _, laser := range g.lasers {
		laser.Update()
	}

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()
		m := NewMeteor()
		g.meteros = append(g.meteros, m)
	}

	for _, meteor := range g.meteros {
		meteor.Update()
	}

	// Verificar colisao no players.
	for _, m := range g.meteros {
		if m.Collider().Intersects(g.player.Collider()) {
			g.Reset()
		}
	}

	// Verificar colisao entre os lasers e os meteoros.
	for i, m := range g.meteros {
		for j, l := range g.lasers {
			if m.Collider().Intersects(l.Collider()) {
				g.meteros = append(g.meteros[:i], g.meteros[i+1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
				g.score += 1
			}
		}
	}

	return nil
}

// Responsavel por desenhar o jogo na tela.
// 60 X por segundo.
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, laser := range g.lasers {
		laser.Draw(screen)
	}

	for _, meteor := range g.meteros {
		meteor.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("Pontos : %d", g.score), assets.FontUi, 20, 100, color.White)
}

// Responsavel por definir o tamanho da tela.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Responsavel por criar um novo laser.
func (g *Game) AddLaser(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) Reset() {
	g.player = NewPlayer(g)
	g.lasers = nil
	g.meteros = nil
	g.meteorSpawnTimer.Reset()
	g.score = 0
}
