package txtimg

import (
	"image"
	"image/color"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func GetDimensions(frames []string) (int, int) {
	rows := strings.Split(frames[0], "\n")
	height := len(rows)
	width := 0

	for _, row := range rows {
		width = max(width, len(strings.Split(row, "")))
	}

	return width, height
}

func BuildGifFrames(frames []string) ([]*image.Paletted, error) {
	return BuildGifFramesWithTick(frames, func(i int) {})
}

func BuildGifFramesWithTick(frames []string, tick func(int)) ([]*image.Paletted, error) {
	var pals []*image.Paletted
	width, height := GetDimensions(frames)

	for i, content := range frames {
		tick(i)

		canvas := NewCanvas(width, height)
		canvas.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
		canvas.Letters(content)

		enc, err := canvas.AsPaletted()

		if err != nil {
			return nil, err
		}

		pals = append(pals, enc)
	}

	return pals, nil
}
