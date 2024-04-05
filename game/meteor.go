package game

import (
	"math/rand"
	"spaceship_go/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Meteor struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewMeteor() *Meteor {
	image := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]
	speed := (rand.Float64() * 13)

	position := Vector{
		X: rand.Float64() * screenWidth,
		Y: -100,
	}

	return &Meteor{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (m *Meteor) Update() {
	m.position.Y += m.speed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// Posição X e Y que a imagem será desenhada na tela do game.
	op.GeoM.Translate(m.position.X, m.position.Y)

	// Desenha a imagem na tela do game.
	screen.DrawImage(m.image, op)
}

func (m *Meteor) Collider() Rect {
	bouds := m.image.Bounds()

	return NewRect(
		m.position.X,
		m.position.Y,
		float64(bouds.Dx()),
		float64(bouds.Dy()),
	)
}
