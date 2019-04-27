package passphoto

import (
	"image"
	"image/color"

	// need jpeg to decode
	_ "image/jpeg"
	"log"
	"os"
)

// GetCanvas, generate a white canvas
func GetCanvas(width, height int) *image.RGBA {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.White)
		}
	}
	return img
}

// Read jpg image
func ReadJpg(path string) image.Image {

	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return m
}
