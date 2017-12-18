package txtimg

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"strings"

	"github.com/golang/freetype/truetype"
	"github.com/minond/txtimg/font/hackregular"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Canvas struct {
	Img        *image.RGBA
	imgHeight  int
	imgWidth   int
	charWidth  int
	charHeight int
}

const (
	defaultFontSize   = 64
	defaultCharWidth  = 44
	defaultCharHeight = 60
)

var (
	monoFont font.Face
)

func init() {
	monoFont = ttf(hackregular.TTF, defaultFontSize)
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

func (c *Canvas) Fill(col color.RGBA) {
	for y := 0; y < c.imgHeight; y++ {
		for x := 0; x < c.imgWidth; x++ {
			c.Img.Set(x, y, col)
		}
	}
}

func (c *Canvas) Write(col color.RGBA, x, y int, str string) {
	point := fixed.Point26_6{
		fixed.Int26_6(x * 62),
		fixed.Int26_6(y * 62),
	}

	d := &font.Drawer{
		Dst:  c.Img,
		Src:  image.NewUniform(col),
		Dot:  point,
		Face: monoFont,
	}

	d.DrawString(str)
}

func (c *Canvas) Letter(x, y int, lttr string) {
	c.Write(color.RGBA{0x00, 0x00, 0x00, 0xff},
		x*c.charWidth, y*c.charHeight+c.charHeight, lttr)
}

func (c *Canvas) Letters(content string) {
	for y, row := range strings.Split(content, "\n") {
		for x, col := range strings.Split(row, "") {
			c.Letter(x, y, col)
		}
	}
}

func (c *Canvas) AsGif() (image.Image, error) {
	buffer := new(bytes.Buffer)
	gif.Encode(buffer, c.Img, &gif.Options{NumColors: 256})
	return gif.Decode(buffer)
}

func (c *Canvas) AsPaletted() (*image.Paletted, error) {
	enc, err := c.AsGif()

	if err != nil {
		return nil, err
	}

	return enc.(*image.Paletted), nil
}

func NewCanvas(w, h int) *Canvas {
	imgWidth := w*defaultCharWidth - defaultCharWidth
	imgHeight := h*defaultCharHeight - defaultCharHeight
	rec := image.Rect(0, 0, imgWidth, imgHeight)
	img := image.NewRGBA(rec)

	return &Canvas{
		Img:        img,
		imgHeight:  imgHeight,
		imgWidth:   imgWidth,
		charWidth:  defaultCharWidth,
		charHeight: defaultCharHeight,
	}
}
