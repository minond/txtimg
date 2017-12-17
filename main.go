package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/math/fixed"
)

const (
	fontSize = 16
)

var (
	monoFont font.Face
)

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

func letter(img *image.RGBA, x, y int, lttr string) {
	write(img, color.RGBA{0x00, 0x00, 0x00, 0xff}, x*fontSize, (y+1)*fontSize, lttr)
}

func canvas(w, h int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, w*fontSize, h*fontSize))
}

func init() {
	monoFont = ttf(gomono.TTF, fontSize)
}

func main() {
	str := strings.TrimSpace(`
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
x                               x
x    o               o          x
x                               x
x              o                x
x                               x
x                               x
x         o                     x
x                               x
x                               x
x                   o           x
x                               x
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
`)

	rows := strings.Split(str, "\n")
	cols := strings.Split(rows[0], "")

	img := canvas(len(cols), len(rows))

	for y, row := range strings.Split(str, "\n") {
		for x, col := range strings.Split(row, "") {
			letter(img, x, y, col)
		}
	}

	handler, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer handler.Close()
	png.Encode(handler, img)
}
