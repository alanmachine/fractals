// Sierpinski triangle
package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"math/rand"
	"time"
)

var (
	triangle = []map[string]float64{
		{"x": 400, "y": 540},
		{"x": 80, "y": 60},
		{"x": 720, "y": 60},
	}

	xRand, yRand float64 = 300, 200
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Sierpinski triangle",
		Bounds: pixel.R(0, 0, 800, 600),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(255, 255, 255)

	imd.Push(pixel.V(xRand, yRand))
	imd.Push(pixel.V(xRand+1, yRand+1))
	imd.Line(1)

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(50, 550), basicAtlas)
	_, _ = fmt.Fprint(basicTxt, "Sierpinski triangle")
	basicTxt.Draw(win, pixel.IM)

	for _, val := range triangle {
		x, y := val["x"], val["y"]
		imd.Push(pixel.V(x, y))
		imd.Push(pixel.V(x+1, y+1))
		imd.Line(1)
	}

	count := 0
	for !win.Closed() {
		count++
		imd.Clear()
		point := triangle[uint8(rand.Intn(3))]
		xRand, yRand = xRand+((point["x"]-xRand)/2), yRand+((point["y"]-yRand)/2)
		drawPoint(xRand, yRand, imd, win)
		drawText(fmt.Sprintln("Iterations: i =", count), imd, win)
		win.Update()
	}
}

func drawPoint(x float64, y float64, imd *imdraw.IMDraw, win *pixelgl.Window) {
	imd.Color = pixel.RGB(255, 255, 255)
	imd.Push(pixel.V(x, y))
	imd.Push(pixel.V(x+1, y+1))
	imd.Line(1)

	imd.Draw(win)
}

func drawText(txt string, imd *imdraw.IMDraw, win *pixelgl.Window) {
	imd.Color = pixel.RGB(0, 0, 0)
	imd.Push(pixel.V(50, 380))
	imd.Push(pixel.V(200, 420))
	imd.Rectangle(0)
	imd.Draw(win)

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(50, 400), basicAtlas)
	_, _ = fmt.Fprint(basicTxt, txt)
	basicTxt.Draw(win, pixel.IM)
}
