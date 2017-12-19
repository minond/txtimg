package main

import (
	"flag"
	"fmt"
	"image/gif"
	"io/ioutil"
	"log"
	"net/http"
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
	var listen string
	var paths []string

	flags := flag.NewFlagSet("txtimg", flag.ExitOnError)
	flags.IntVar(&delay, "delay", 25, "Delay between gif frames.")
	flags.StringVar(&listen, "listen", "", "Host and port to bind server to.")

	usage := flags.Usage
	flags.Usage = func() {
		usage()
		fmt.Println("  <frames>* []string")
		fmt.Println("        Path to frame files.")
	}

	flags.Parse(os.Args[1:])

	// server
	if listen != "" {
		fmt.Printf("Setting up server on %s\n", listen)
		http.ListenAndServe(listen, txtimg.Service())
		return
	}

	paths = flags.Args()

	if len(paths) == 0 {
		flags.Usage()
		os.Exit(2)
	}

	out := &gif.GIF{}
	frames := files(paths...)
	handler, err := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer handler.Close()

	if err != nil {
		log.Fatalf("Error opening output file: %v", err)
	}

	images, err := txtimg.BuildGifFramesWithTick(frames, func(i int) {
		fmt.Printf("%3.0f%% - Encoding %s\n", float64(i+1)/float64(len(frames))*100, paths[i])
	})

	if err != nil {
		log.Fatalf("Error encoding: %v", err)
	}

	out.Image = append(out.Image, images...)

	for i := 0; i < len(images); i++ {
		out.Delay = append(out.Delay, delay)
	}

	gif.EncodeAll(handler, out)
	fmt.Printf("Done - Saved to out.gif with a delay of %d between frames\n", delay)
}
