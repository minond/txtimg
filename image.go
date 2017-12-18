package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/golang/freetype/truetype"
	"github.com/minond/txtimg/font/hackregular"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	fontSize   = 64
	charWidth  = 44
	charHeight = 60
)

var (
	monoFont font.Face
)

func init() {
	monoFont = ttf(hackregular.TTF, fontSize)
}

func ttf(data []byte, size int) font.Face {
	f, err := truetype.Parse(data)

	if err != nil {
		panic(fmt.Sprintf("Error parsing TTF: %v", err))
	}

	return truetype.NewFace(f, &truetype.Options{
		SubPixelsX: 64,
		SubPixelsY: 64,
		Hinting:    font.HintingFull,
		Size:       float64(size),
	})
}

func write(img *image.RGBA, col color.RGBA, x, y int, str string) {
	point := fixed.Point26_6{
		fixed.Int26_6(x * 62),
		fixed.Int26_6(y * 62),
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Dot:  point,
		Face: monoFont,
	}

	d.DrawString(str)
}

func dot(img *image.RGBA, x, y int, col color.RGBA) {
	img.Set(x, y, col)
}

func fill(img *image.RGBA, w, h int, col color.RGBA) {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			dot(img, x, y, col)
		}
	}
}

func letter(img *image.RGBA, x, y int, lttr string) {
	write(img, color.RGBA{0x00, 0x00, 0x00, 0xff},
		x*charWidth, y*charHeight+charHeight, lttr)
}

func canvas(w, h int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, w*charWidth-charWidth, h*charHeight-charHeight))
}
