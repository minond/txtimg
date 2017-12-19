package main

import (
	"flag"
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
	var delay int
	var paths []string

	flags := flag.NewFlagSet("txtimg", flag.ExitOnError)
	flags.IntVar(&delay, "delay", 25, "Delay between gif frames.")

	usage := flags.Usage
	flags.Usage = func() {
		usage()
		fmt.Println("  <paths>* []string")
		fmt.Println("        Path to frame files.")
	}

	flags.Parse(os.Args[1:])
	paths = flags.Args()

	if len(paths) == 0 {
		flags.Usage()
		os.Exit(2)
	}

	out := &gif.GIF{}
	frames := files(paths...)
	handler, err := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatalf("Error opening output file: %v", err)
	}

	defer handler.Close()

	images, err := txtimg.BuildGifFramesWithTick(frames, func(i int) {
		fmt.Printf("Encoding %s (%02d/%02d)\n", paths[i], i+1, len(frames))
	})

	if err != nil {
		log.Fatalf("Error encoding: %v", err)
	}

	out.Image = append(out.Image, images...)

	for i := 0; i < len(images); i++ {
		out.Delay = append(out.Delay, delay)
	}

	gif.EncodeAll(handler, out)
	fmt.Printf("Saved to out.gif with a delay of %d between frames.\n", delay)
}
