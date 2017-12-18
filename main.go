package main

import (
	"fmt"
	"image/png"
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
	fmt.Println("Usage: go run main.go in.txt out.png")
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
