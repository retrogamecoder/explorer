package resources

import (
	"sync"

	"github.com/faiface/pixel"
)

type Manager interface {
	GetImage(ImageID) (pixel.Picture, bool)
}

type Config struct {
	Images map[ImageID]string
}

func NewManager(cfg Config) (Manager, error) {
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
