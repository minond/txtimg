package main

import (
	"fmt"
	"image/gif"
	"io/ioutil"
	"log"
	"os"

	"github.com/minond/txtimg"
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

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	out := &gif.GIF{}
	frames := files(os.Args[1:]...)
	handler, err := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatalf("Error opening output file: %v", err)
	}

	defer handler.Close()

	images, err := txtimg.BuildGifFramesWithTick(frames, func(i int) {
		fmt.Printf("Generating %d out of %d frames.\n", i+1, len(frames))
	})

	if err != nil {
		log.Fatalf("Error encoding: %v", err)
	}

	out.Image = append(out.Image, images...)

	for i := 0; i < len(images); i++ {
		out.Delay = append(out.Delay, 25)
	}

	gif.EncodeAll(handler, out)
	fmt.Println("Saved to out.gif")
}
