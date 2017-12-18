package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io/ioutil"
	"os"
	"strings"
)

func file(path string) string {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		os.Exit(2)
	}

	return string(content)
}

func usage() {
	fmt.Println("Usage: go run main.go [in.txt]*")
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
		// str = file(os.Args[1])
	}

	out := &gif.GIF{}
	files := os.Args[1:]
	var contents []string

	for _, v := range files {
		contents = append(contents, file(v))
	}

	rows := strings.Split(contents[0], "\n")
	rowsLen := len(rows)
	colsLen := 0

	// Fix widest row
	for _, row := range rows {
		colsLen = max(colsLen, len(strings.Split(row, "")))
	}

	for i, content := range contents {
		fmt.Printf("Generating %d out of %d frames...", i+1, len(contents))
		img := canvas(colsLen, rowsLen)

		// Fill background
		fill(img, colsLen*charWidth, rowsLen*charHeight,
			color.RGBA{0xff, 0xff, 0xff, 0xff})

		// Draw map
		for y, row := range strings.Split(content, "\n") {
			for x, col := range strings.Split(row, "") {
				letter(img, x, y, col)
			}
		}

		buffer := new(bytes.Buffer)
		gif.Encode(buffer, img, &gif.Options{NumColors: 256})
		encoded, _ := gif.Decode(buffer)

		out.Image = append(out.Image, encoded.(*image.Paletted))
		out.Delay = append(out.Delay, 25)

		fmt.Printf(" ok!\n")
	}

	handler, _ := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer handler.Close()

	gif.EncodeAll(handler, out)
}
