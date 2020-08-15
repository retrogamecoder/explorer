package view

import (
	"sync"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/retrogamecoder/explorer/internal/model"
	"golang.org/x/image/colornames"
)

const (
	pixelWidth  = 32
	pixelHeight = 32

	playerSpriteXOffset = 7
	playerSpriteYOffset = 5
)

// An ImageID is the unique identifier for an image that is loaded from disk.
type ImageID string

// A View is something that can render a World.
type View interface {
	Render(world *model.World, window *pixelgl.Window)
}

// A ViewConfig is the initialization settings for a View.
type ViewConfig struct {
	Images            map[ImageID]string
	PlayerSpriteImage ImageID
}

// NewView constructs a default View object based on the config passed.
func NewView(cfg ViewConfig) (View, error) {
	images := map[ImageID]pixel.Picture{}
	var wg sync.WaitGroup
	wg.Add(len(cfg.Images))

	for imageID, path := range cfg.Images {
		go func(imageID ImageID, path string) {
			img, _ := loadImage(path)
			images[imageID] = img
			wg.Done()
		}(imageID, path)
	}

	wg.Wait()

	playerSprite := pixel.NewSprite(images[cfg.PlayerSpriteImage], pixel.R(
		pixelWidth*playerSpriteXOffset,
		pixelHeight*playerSpriteYOffset,
		pixelWidth*(playerSpriteXOffset+1),
		pixelHeight*(playerSpriteYOffset+1),
	))

	return &view{
		images,
		playerSprite,
	}, nil
}

type view struct {
	images       map[ImageID]pixel.Picture
	playerSprite *pixel.Sprite
}

func (v *view) Render(world *model.World, window *pixelgl.Window) {
	window.Clear(colornames.Black)
	v.playerSprite.Draw(window, pixel.IM.Moved(world.Player.Pos))
}
