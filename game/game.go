package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player *Player
	lasers []*Laser
}

func NewGame() *Game {
	g := &Game{}
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

	return nil
}

// Responsavel por desenhar o jogo na tela.
// 60 X por segundo.
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, laser := range g.lasers {
		laser.Draw(screen)
	}
}

// Responsavel por definir o tamanho da tela.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Responsavel por criar um novo laser.
func (g *Game) AddLaser(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}
