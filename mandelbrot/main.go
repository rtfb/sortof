// Mandelbrot generation code starts off from gopl.io/ch3/mandelbrot/.
package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"math/cmplx"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	width, height = 1024, 1024
)

// zframe represents a normalized zoom frame. x and y are the pair of
// normalised coordinates of the lower left corner of the zoom square, and size
// is the actual size of the square. A full zoom out is therefore a x=0, y=0,
// size=1. A 10x zoom-in from the max will leave x,y at zero, and will set size
// to 0.1. Another 10x zoom-in will set size to 0.01, etc.
type zframe struct {
	x    float64
	y    float64
	size float64
}

// mframe represents a zoom frame converted to the coordinates in numbers that
// make sense for mandelbrot rendering. A full zoom out is <-2, -2>, <2, 2>.
type mframe struct {
	xmin float64
	ymin float64
	xmax float64
	ymax float64
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	currZoom := zframe{
		size: 1.0,
	}
	zoomStack := []zframe{currZoom}
	zoomSP := 0
	fmt.Printf("currZoom: %#v\n", currZoom)

	var zoom [2]pixel.Vec
	zoomingIn := false
	zoomPos := 0
	zoomIn := false
	zoomOut := false
	sprite := repaint(currZoom)

	win.Clear(colornames.Skyblue)
	for !win.Closed() {
		if win.JustPressed(pixelgl.KeyEscape) {
			break
		}
		if win.JustPressed(pixelgl.KeyO) {
			zoomOut = true
		}
		if win.JustPressed(pixelgl.MouseButtonLeft) {
			fmt.Printf("%#v\n", win.MousePosition())
			if !zoomingIn {
				zoomingIn = true
			}
			if zoomingIn {
				zoom[zoomPos] = win.MousePosition()
				zoomPos++
				if zoomPos > 1 {
					zoomingIn = false
					zoomPos = 0
					fmt.Printf("zoom: p1=%v, p2=%v\n", zoom[0], zoom[1])
					zoomIn = true
				}
			}
		}
		if zoomOut {
			zoomSP--
			if zoomSP < 0 {
				zoomSP = 0
			}
			currZoom = zoomStack[zoomSP]
			zoomOut = false
			sprite = repaint(currZoom)
		}
		if zoomIn {
			currZoom = calcZoom(currZoom, zoom)
			fmt.Printf("currZoom: %#v\n", currZoom)
			zoomStack = append(zoomStack, currZoom)
			zoomSP++
			sprite = repaint(currZoom)
			zoomIn = false
		}
		sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}

func calcZoom(currZoom zframe, zoom [2]pixel.Vec) zframe {
	x1 := zoom[0].X / width
	x2 := zoom[1].X / width
	y1 := zoom[0].Y / height
	y2 := zoom[1].Y / height
	dx := math.Abs(x1 - x2)
	dy := math.Abs(y1 - y2)
	newSize := dx
	if dy > newSize {
		newSize = dy
	}
	midx := (x1 + x2) / 2
	midy := (y1 + y2) / 2
	return zframe{
		x:    currZoom.x + (midx-newSize/2)*currZoom.size,
		y:    currZoom.y + (midy-newSize/2)*currZoom.size,
		size: newSize * currZoom.size,
	}
}

func zframeToM(z zframe) mframe {
	xmin := (z.x * 4) - 2
	ymin := (z.y * 4) - 2
	return mframe{
		xmin: xmin,
		xmax: xmin + z.size*4,
		ymin: ymin + z.size*4,
		ymax: ymin,
	}
}

func repaint(zoom zframe) *pixel.Sprite {
	mzoom := zframeToM(zoom)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(mzoom.ymax-mzoom.ymin) + mzoom.ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(mzoom.xmax-mzoom.xmin) + mzoom.xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}

	pic := pixel.PictureDataFromImage(img)
	return pixel.NewSprite(pic, pic.Bounds())
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			pixel := 255 - contrast*n
			var r, g, b uint8
			if pixel < 80 {
				r = (85 - pixel) + 170
				g, b = 0, 0
			} else if pixel < 160 {
				g = (170 - pixel) + 95
				r, b = 0, 0
			} else {
				b = 100 - pixel
				r, g = 0, 0
			}
			return color.RGBA{
				R: r,
				G: g,
				B: b,
				A: 255,
			}
		}
	}
	return color.Black
}
