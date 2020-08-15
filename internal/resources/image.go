package resources

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
)

func loadImage(path string) (pixel.Picture, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
