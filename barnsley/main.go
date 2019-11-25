// Barnsley fern
package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	pixelgl.Run(run)
}

func run() {
	var x, y float64

	cfg := pixelgl.WindowConfig{
		Title:  "Barnsley fern",
		Bounds: pixel.R(0, 0, 800, 600),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(50, 550), basicAtlas)
	_, _ = fmt.Fprint(basicTxt, "Barnsley fern")
	basicTxt.Draw(win, pixel.IM)

	count := 0
	for !win.Closed() {
		count++
		imd.Clear()
		randInt := uint8(rand.Intn(100)) + 1
		x, y = affineTrans(randInt, x, y)
		imd.Color = setColor(randInt)

		drawPoint(x*60+400, y*50+50, imd, win)
		drawText(fmt.Sprintln("Iterations: i =", count), imd, win)
		win.Update()
	}
}

func affineTrans(randInt uint8, xIn float64, yIn float64) (xOut, yOut float64) {
	switch {
	case randInt <= 2: // 1%
		xOut = 0
		yOut = 0.16 * yIn
	case randInt <= 86: // 85 %
		xOut = 0.85*xIn + 0.04*yIn
		yOut = -0.04*xIn + 0.85*yIn + 1.6
	case randInt <= 93: // 7%
		xOut = 0.2*xIn - 0.26*yIn
		yOut = 0.23*xIn + 0.22*yIn + 1.6
	default: // 7%
		xOut = -0.15*xIn + 0.28*yIn
		yOut = 0.26*xIn + 0.24*yIn + 0.44
	}

	return xOut, yOut
}

func setColor(randInt uint8) color.RGBA {
	if randInt <= 86 {
		return color.RGBA{R: 0, G: 221, B: 0, A: 255}
	}

	return color.RGBA{R: 0, G: 164, B: 0, A: 255}
}

func drawPoint(xIn float64, yIn float64, imd *imdraw.IMDraw, win *pixelgl.Window) {
	imd.Push(pixel.V(xIn, yIn))
	imd.Push(pixel.V(xIn+1, yIn+1))
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
