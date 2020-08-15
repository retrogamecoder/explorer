package view

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/retrogamecoder/explorer/internal/model"
	"github.com/retrogamecoder/explorer/internal/resources"
	"golang.org/x/image/colornames"
)

const (
	pixelWidth  = 32
	pixelHeight = 32

	playerSpriteXOffset = 7
	playerSpriteYOffset = 5
)

// A View is something that can render a World.
type View interface {
	Render(world *model.World, window *pixelgl.Window)
}

// A ViewConfig is the initialization settings for a View.
type ViewConfig struct {
	ResourceManager   resources.Manager
	PlayerSpriteImage resources.ImageID
}

// NewView constructs a default View object based on the config passed.
func NewView(cfg ViewConfig) (View, error) {
	playerImg, ok := cfg.ResourceManager.GetImage(cfg.PlayerSpriteImage)
	if !ok {
		return nil, fmt.Errorf("Player sprite ID %s does not exist", cfg.PlayerSpriteImage)
	}

	playerSprite := pixel.NewSprite(playerImg, pixel.R(
		pixelWidth*playerSpriteXOffset,
		pixelHeight*playerSpriteYOffset,
		pixelWidth*(playerSpriteXOffset+1),
		pixelHeight*(playerSpriteYOffset+1),
	))

	return &view{
		cfg.ResourceManager,
		playerSprite,
	}, nil
}

type view struct {
	res          resources.Manager
	playerSprite *pixel.Sprite
}

func (v *view) Render(world *model.World, window *pixelgl.Window) {
	window.Clear(colornames.Black)
	v.playerSprite.Draw(window, pixel.IM.Moved(world.Player.Pos))
}
