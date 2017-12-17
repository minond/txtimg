package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
	"strings"

	"github.com/golang/freetype/truetype"
	"github.com/minond/txtimg/hackregular"
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
	write(img, color.RGBA{0x00, 0x00, 0x00, 0xff}, x*charWidth, y*charHeight+charHeight, lttr)
}

func canvas(w, h int) *image.RGBA {
	return image.NewRGBA(image.Rect(0, 0, w*charWidth-charWidth, h*charHeight))
}

func stdio() string {
	scanner := bufio.NewScanner(os.Stdin)
	content := ""

	for scanner.Scan() {
		content += "\n" + scanner.Text()
	}

	return content
}

func file(path string) string {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		os.Exit(2)
	}

	return string(content)
}

func usage() {
	fmt.Println("Usage: go run main.go out.png < in.txt")
	fmt.Println("       go run main.go in.txt out.png")
}

func init() {
	monoFont = ttf(hackregular.TTF, fontSize)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var str string
	var out string

	if len(os.Args) > 2 {
		str = file(os.Args[1])
		out = os.Args[2]
	} else if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
		str = stdio()
		out = os.Args[1]
	} else {
		usage()
		os.Exit(2)
	}

	rows := strings.Split(str, "\n")
	colsLen := 0

	// Fix widest row
	for _, row := range rows {
		colsLen = max(colsLen, len(strings.Split(row, "")))
	}

	img := canvas(colsLen, len(rows))

	for y, row := range strings.Split(str, "\n") {
		for x, col := range strings.Split(row, "") {
			letter(img, x, y, col)
		}
	}

	handler, _ := os.OpenFile(out, os.O_WRONLY|os.O_CREATE, 0600)
	defer handler.Close()
	png.Encode(handler, img)
}
