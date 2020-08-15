package controller

import (
	"time"

	"github.com/faiface/pixel/pixelgl"
	"github.com/retrogamecoder/explorer/internal/model"
)

type Controller interface {
	Update(world *model.World, window *pixelgl.Window)
}

func NewController() Controller {
	return &controller{
		time.Now(),
	}
}

type controller struct {
	lastUpdate time.Time
}

func (c *controller) Update(world *model.World, window *pixelgl.Window) {
	dt := time.Since(c.lastUpdate).Seconds()
	c.lastUpdate = time.Now()

	if window.Pressed(pixelgl.KeyLeft) {
		world.Player.Pos.X -= world.Player.MoveSpeed * dt
	}
	if window.Pressed(pixelgl.KeyRight) {
		world.Player.Pos.X += world.Player.MoveSpeed * dt
	}
	if window.Pressed(pixelgl.KeyUp) {
		world.Player.Pos.Y += world.Player.MoveSpeed * dt
	}
	if window.Pressed(pixelgl.KeyDown) {
		world.Player.Pos.Y -= world.Player.MoveSpeed * dt
	}
}
