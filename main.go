package main

import (
	"fmt"
	"image/color"
	"image/gif"
	"io/ioutil"
	"log"
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

func files(paths ...string) []string {
	var contents []string

	for _, path := range paths {
		contents = append(contents, file(path))
	}

	return contents
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

func getDimensions(frame string) (int, int) {
	rows := strings.Split(frame, "\n")
	height := len(rows)
	width := 0

	for _, row := range rows {
		width = max(width, len(strings.Split(row, "")))
	}

	return width, height
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	out := &gif.GIF{}
	frames := files(os.Args[1:]...)
	width, height := getDimensions(frames[0])
	handler, err := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatalf("Error opening output file: %v", err)
	}

	defer handler.Close()

	for i, content := range frames {
		fmt.Printf("Generating %d out of %d frames...", i+1, len(frames))

		canvas := NewCanvas(width, height)
		canvas.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
		canvas.Letters(content)

		enc, err := canvas.asPaletted()

		if err != nil {
			log.Fatalf("Error encoding frame #%d: %v", i, err)
		}

		out.Image = append(out.Image, enc)
		out.Delay = append(out.Delay, 25)

		fmt.Printf(" ok!\n")
	}

	gif.EncodeAll(handler, out)
	fmt.Println("Saved to out.gif")
}
