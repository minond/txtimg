package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/math/fixed"
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
		Face: ttf(gomono.TTF, 24),
	}

	d.DrawString(str)
}

// Gif WIP:
//
// 	var images []*image.Paletted
// 	var delays []int
//
// 	var palette = []color.Color{
// 		color.RGBA{0x00, 0x00, 0x00, 0xFF},
// 		color.RGBA{0x00, 0x00, 0xFF, 0xFF},
// 		color.RGBA{0x00, 0xFF, 0x00, 0xFF},
// 		color.RGBA{0x00, 0xFF, 0xFF, 0xFF},
// 		color.RGBA{0xFF, 0x00, 0x00, 0xFF},
// 		color.RGBA{0xFF, 0x00, 0xFF, 0xFF},
// 		color.RGBA{0xFF, 0xFF, 0x00, 0xFF},
// 		color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
// 	}
//
// 	var w, h int = 240, 240
//
// 	for i := 0; i < 100; i++ {
// 		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
// 		images = append(images, img)
// 		delays = append(delays, 3)
//
// 		img.Set(40, 40+i, color.RGBA{uint8(55), uint8(255), uint8(55), 255})
// 		img.Set(40, 41+i, color.RGBA{uint8(55), uint8(255), uint8(55), 255})
// 		img.Set(40, 42+i, color.RGBA{uint8(55), uint8(255), uint8(55), 255})
// 		img.Set(40, 43+i, color.RGBA{uint8(55), uint8(255), uint8(55), 255})
// 		img.Set(40, 44+i, color.RGBA{uint8(55), uint8(255), uint8(55), 255})
// 	}
//
// 	f, _ := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
// 	defer f.Close()
//
// 	gif.EncodeAll(f, &gif.GIF{
// 		Image: images,
// 		Delay: delays,
// 	})
func main() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	write(img, color.RGBA{0x00, 0x00, 0x00, 0xff}, 20, 20, "xo")
	write(img, color.RGBA{0x00, 0x00, 0x00, 0xff}, 20, 44, "xo")

	handler, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer handler.Close()

	png.Encode(handler, img)
}
