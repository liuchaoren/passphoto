package passphoto

import (
	"image"
	"log"
)

func MergePhoto(path string) {

	img := ReadJpg(path)

	bounds := img.Bounds()
	imgWidth := bounds.Dx()
	imgHeight := bounds.Dy()

	if imgWidth != imgHeight {
		log.Fatal("image must be a square")
	}
	canvasWidth, canvasHeight := canvasSize(imgWidth)
	canvas := GetCanvas(canvasWidth, canvasHeight)

	updateCanvas := populatePixel(img, canvas)
	WriteJpg(updateCanvas, "output.png")
}

func canvasSize(imgSize int) (int, int) {
	canvasWidth := imgSize * 3
	canvasHeight := imgSize * 2
	return canvasWidth, canvasHeight
}

func populatePixel(img image.Image, canvas *image.RGBA) *image.RGBA {
	imgsize := img.Bounds().Dx()

	widthOffset := imgsize / 3
	heightOffset := imgsize / 2
	for i := 0; i < imgsize; i++ {
		for j := 0; j < imgsize; j++ {
			xindex := i + widthOffset
			yindex := j + heightOffset
			canvas.Set(xindex, yindex, img.At(i, j))
		}
	}

	newWidthOffset := widthOffset*2 + imgsize
	for i := 0; i < imgsize; i++ {
		for j := 0; j < imgsize; j++ {
			xindex := i + newWidthOffset
			yindex := j + heightOffset
			canvas.Set(xindex, yindex, img.At(i, j))
		}
	}
	return canvas
}
