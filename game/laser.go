package game

import (
	"spaceship_go/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite

	halfW := float64(image.Bounds().Dx()) / 2 // Dividindo a largura da imagem por 2.
	halfH := float64(image.Bounds().Dy()) / 2 // Dividindo a altura da imagem por 2.

	position.X -= halfW
	position.Y -= halfH

	return &Laser{
		image:    image,
		position: position,
	}
}

func (l *Laser) Update() {
	speed := 7.0

	l.position.Y += -speed
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Posição X e Y que a imagem será desenhada na tela do game.
	op.GeoM.Translate(l.position.X, l.position.Y)

	// Desenha a imagem na tela do game.
	screen.DrawImage(l.image, op)
}
