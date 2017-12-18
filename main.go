package main

import (
	"fmt"
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
	fmt.Println("Usage: go run main.go in.txt out.gif")
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
	} else {
		usage()
		os.Exit(2)
	}

	rows := strings.Split(str, "\n")
	rowsLen := len(rows)
	colsLen := 0

	// Fix widest row
	for _, row := range rows {
		colsLen = max(colsLen, len(strings.Split(row, "")))
	}

	img := canvas(colsLen, rowsLen)

	// Fill background
	fill(img, colsLen*charWidth, rowsLen*charHeight,
		color.RGBA{0xff, 0xff, 0xff, 0xff})

	for y, row := range strings.Split(str, "\n") {
		for x, col := range strings.Split(row, "") {
			letter(img, x, y, col)
		}
	}

	handler, _ := os.OpenFile(out, os.O_WRONLY|os.O_CREATE, 0600)
	opt := gif.Options{NumColors: 256}

	defer handler.Close()
	gif.Encode(handler, img, &opt)
}
