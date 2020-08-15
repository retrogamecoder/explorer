package resources

import (
	"sync"

	"github.com/faiface/pixel"
	"github.com/lafriks/go-tiled"
)

type Manager interface {
	GetImage(ImageID) (pixel.Picture, bool)
	GetMap(MapID) (*tiled.Map, bool)
}

// A Config is the set of assets that need to be loaded.
type Config struct {
	Images            map[ImageID]string
	Tileset           map[TilesetID]string
	PlayerSpriteImage ImageID
}

func NewManager(cfg Config) (Manager, error) {
	images := map[ImageID]pixel.Picture{}
	var wg sync.WaitGroup
	wg.Add(len(cfg.Images))

	for imageID, path := range cfg.Images {
		go func(imageID ImageID, path string) {
			// TODO(rob): Handle errors
			img, _ := loadImage(path)
			images[imageID] = img
			wg.Done()
		}(imageID, path)
	}

	wg.Wait()

	return &manager{
		images,
	}, nil
}

type manager struct {
	images map[ImageID]pixel.Picture
}

func (m *manager) GetImage(id ImageID) (pixel.Picture, bool) {
	img, ok := m.images[id]
	return img, ok
}

func (m *manager) GetMap(id MapID) (*tiled.Map, bool) {
	return nil, false
}
