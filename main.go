package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
)

func main() {
	path, err := filepath.Abs(os.Args[1])

	if err != nil {
		panic(err)
	}

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	i, _, err := image.Decode(file)

	if err != nil {
		panic(err)
	}

	bg := i.At(0, 0)
	out := image.NewRGBA(i.Bounds())

	for x := 0; x < i.Bounds().Max.X; x++ {
		for y := 0; y < i.Bounds().Max.Y; y++ {
			if i.At(x, y) == bg {
				out.Set(x, y, color.Transparent)
			} else {
				out.Set(x, y, i.At(x, y))
			}
		}
	}

	png.Encode(os.Stdout, out)
}
