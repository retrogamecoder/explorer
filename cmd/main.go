package main

import (
	"log"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/retrogamecoder/explorer/internal/controller"
	"github.com/retrogamecoder/explorer/internal/model"
	"github.com/retrogamecoder/explorer/internal/view"
)

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

	v, err := view.NewView(view.ViewConfig{
		Images: map[view.ImageID]string{
			"base": "sprites/sample.png",
		},
		PlayerSpriteImage: "base",
	})

	world := &model.World{
		Player: &model.Player{
			Entity: model.Entity{
				Pos:       w.Bounds().Center(),
				MoveSpeed: 100.0,
			},
		},
	}

	c := controller.NewController()

	for !w.Closed() {
		c.Update(world, w)
		v.Render(world, w)
		w.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
