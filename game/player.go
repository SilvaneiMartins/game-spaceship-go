package game

import (
	"spaceship_go/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image             *ebiten.Image
	position          Vector
	game              *Game
	laserLoadingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	bouds := image.Bounds()
	halfW := float64(bouds.Dx()) / 2

	position := Vector{
		X: (screenWidth / 2) - halfW,
		Y: 500,
	}

	return &Player{
		image:             image,
		game:              game,
		position:          position,
		laserLoadingTimer: NewTimer(12),
	}
}

func (p *Player) Update() {
	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	}

	p.laserLoadingTimer.Update()

	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTimer.IsReady() {
		p.laserLoadingTimer.Reset()

		bouds := p.image.Bounds()
		halfW := float64(bouds.Dx()) / 2
		halfH := float64(bouds.Dy()) / 2

		spawnPos := Vector{
			X: p.position.X + halfW,
			Y: p.position.Y - halfH/2,
		}

		laser := NewLaser(spawnPos)
		p.game.AddLaser(laser)
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Posição X e Y que a imagem será desenhada na tela do game.
	op.GeoM.Translate(p.position.X, p.position.Y)

	// Desenha a imagem na tela do game.
	screen.DrawImage(p.image, op)
}

func (p *Player) Collider() Rect {
	bouds := p.image.Bounds()

	return NewRect(
		p.position.X,
		p.position.Y,
		float64(bouds.Dx()),
		float64(bouds.Dy()),
	)
}
