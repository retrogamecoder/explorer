package main

import (
	"image"
	"log"
	"os"
	"time"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func loadImage(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "RetroGameCoder - Explorer",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	w, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatalf("Could not create window: %s", err)
	}

	img, err := loadImage("sample.png")
	if err != nil {
		log.Fatalf("Could not load image: %s", err)
	}

	sprite := pixel.NewSprite(img, pixel.R(7*32, 5*32, 8*32, 6*32))

	v := w.Bounds().Center()
	moveSpeed := 100.0
	last := time.Now()

	for !w.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		if w.Pressed(pixelgl.KeyLeft) {
			v.X -= moveSpeed * dt
		}
		if w.Pressed(pixelgl.KeyRight) {
			v.X += moveSpeed * dt
		}
		if w.Pressed(pixelgl.KeyUp) {
			v.Y += moveSpeed * dt
		}
		if w.Pressed(pixelgl.KeyDown) {
			v.Y -= moveSpeed * dt
		}

		w.Clear(colornames.Black)
		sprite.Draw(w, pixel.IM.Moved(v))

		w.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
