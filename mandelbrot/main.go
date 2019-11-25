// Mandelbrot set
package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"log"
	"math"
	"math/cmplx"
	"time"
)

var x, y, scale = 0, 0, 1

const (
	width, height = 640, 640
)

func main() {
	pixelgl.Run(run)
}

func run() {
	var speed int
	cfg := pixelgl.WindowConfig{
		Title: "Mandelbrot Set",
		Bounds: pixel.R(0, 0, width, height),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}
	imd := imdraw.New(nil)
	tick := time.Tick(time.Second / 120)
	start := time.Now()
	i := 0.0
	for !win.Closed() {
		if i >= math.MaxFloat64 {
			i = 0.0
			start = time.Now()
		}
		i++

		if win.MouseScroll().Y != 0 {
			switch {
			case scale > 0 && scale < 1000: {
				speed = 10
			}
			case scale > 1000 && scale < 10000: {
				speed = 100
			}
			case scale > 10000 && scale < 100000: {
				speed = 1000
			}
			default:
				speed = 10000
			}
			scale += int(win.MouseScroll().Y) * speed
			x = int(win.MousePosition().X)
			y = height - int(win.MousePosition().Y)
		}

		if win.JustPressed(pixelgl.MouseButtonLeft) {
			x, y, scale = 0, 0, 1
		}

		imd.Color = pixel.RGB(0, 0, 0)
		imd.Push(pixel.V(0, 0))
		imd.Push(pixel.V(width, height))
		imd.Rectangle(0)
		imd.Draw(win)

		imd.Clear()
		draw(imd, win)

		timer := time.Since(start).Seconds()

		mouseY := int(win.MousePosition().Y)
		if mouseY == 0 {
			mouseY = 640
		}

		atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
		txt := text.New(pixel.V(40, 600), atlas)
		txt.Clear()
		_, _ = fmt.Fprint(txt, "Mandelbrot Set\n\n")
		_, _ = fmt.Fprint(txt, "FPS: ", math.Round(i/timer), "\n")
		_, _ = fmt.Fprint(txt, "Mouse: ", int(win.MousePosition().X), height - mouseY, "\n")
		_, _ = fmt.Fprint(txt, "Scale: x", scale)
		txt.Draw(win, pixel.IM)

		win.Update()
		<-tick
	}
}

func draw(imd *imdraw.IMDraw, win *pixelgl.Window) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, 2, 2
	)
	scaleWidth, scaleHeight := width*scale, height*scale
	offsetX, offsetY := x*scale, y*scale
	for py := 0; py < height; py++ {
		y := float64(py+offsetY)/float64(scaleHeight)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px+offsetX)/float64(scaleWidth)*(xmax-xmin) + xmin
			z := complex(x, y)
			imd.Color = mandelbrot(z)
			imd.Push(pixel.V(float64(px), height - float64(py)))
			imd.Push(pixel.V(float64(px+1), height - float64(py+1)))
			imd.Rectangle(0)
		}
	}
	imd.Draw(win)
}

func mandelbrot(z complex128) color.RGBA {
	const (
		iterations = 240
		contrast   = 15
	)
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{R: 0 + contrast*n, G: 150 + contrast*n/2, B: 255 - contrast*n, A: 255}
		}
	}
	return color.RGBA{R: 0, G: 0, B: 0, A: 0}
}
