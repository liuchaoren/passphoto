package passphoto

import (
	"image"
	"image/color"
	"image/jpeg"
	"path/filepath"

	// need jpeg to decode
	_ "image/jpeg"
	"log"
	"os"

	"github.com/liuchaoren/passphoto/common"
)

const (
	cropConfigFileName string = "crop.json"
)

// GetCanvas generates a white rectangle canvas
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

// ReadJpg reads a jpg image
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

// WriteJpg writes an image to jpg file
func WriteJpg(img image.Image, path string, quality int) {
	imagePath := filepath.Join(common.ExportFolder(), path)

	f, err := os.Create(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	jpeg.Encode(f, img, &jpeg.Options{Quality: quality})
}

// ImageSize return the (widht, and height
func ImageSize(img image.Image) (int, int) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	return width, height
}

// caculate how to cut the four edges
func calculateHowToCrop(originalHeadTop, originalHeadBottom,
	originalCenterFromLeft float64, originalWidth, originalHeight int) (float64, float64, float64, float64) {

	cropConfig := common.ReadConfig(cropConfigFileName)
	headTop := cropConfig["head_top"].(float64)
	headBottom := cropConfig["head_bottom"].(float64)

	cutTop := (originalHeadTop*headBottom - originalHeadBottom*headTop) / (headBottom - headTop)
	cutBottom := (originalHeadTop + (headTop-1)*cutTop) / headTop

	squareSideLength := int((cutBottom - cutTop) * float64(originalHeight))
	cutLeft := originalCenterFromLeft - (float64(squareSideLength) / float64(2) / float64(originalWidth))
	cutRight := originalCenterFromLeft + (float64(squareSideLength) / float64(2) / float64(originalWidth))
	return cutLeft, cutRight, cutTop, cutBottom
}

// Crop crops image given the vertical position of head top, head bottom
// and horiental position of head cener
func Crop(img image.Image, originalHeadTop, orginalHeadBottom,
	originalCenterFromLeft float64) image.Image {
	width, height := ImageSize(img)
	cutLeft, cutRight, cutTop, cutBottom := calculateHowToCrop(originalHeadTop,
		orginalHeadBottom, originalCenterFromLeft, width, height)
	return crop(img, cutLeft, cutRight, cutTop, cutBottom)
}

// crop acutall does the crop and returns the cropped image
func crop(img image.Image, cutLeft, cutRight,
	cutTop, cutBottom float64) image.Image {
	w, h := ImageSize(img)
	imgRGBA, ok := img.(*image.RGBA)
	if !ok {
		log.Fatal("cannot convert to RGBA image")
	}

	leftPixel := int(float64(w) * cutLeft)
	rightPixel := int(float64(w) * cutRight)
	topPixel := int(float64(h) * cutTop)
	bottomPixel := int(float64(h) * cutBottom)

	tmpWidth := rightPixel - leftPixel
	tmpHeight := bottomPixel - topPixel

	var squareSideLength int

	if tmpWidth > tmpHeight {
		squareSideLength = tmpWidth
	} else {
		squareSideLength = tmpHeight
	}

	rightPixel = leftPixel + squareSideLength
	bottomPixel = topPixel + squareSideLength

	upLeft := image.Point{leftPixel, topPixel}
	downRight := image.Point{rightPixel, bottomPixel}
	subImage := imgRGBA.SubImage(image.Rectangle{upLeft, downRight})
	return subImage
}

// OutputPhotoForPrint prepares for print on a 4x6 canvas
func OutputPhotoForPrint(path string) {

	img := ReadJpg(path)

	w, h := ImageSize(img)
	if w != h {
		log.Fatal("image must be a square")
	}
	canvasWidth, canvasHeight := canvasSize(w)
	canvas := GetCanvas(canvasWidth, canvasHeight)

	updateCanvas := populatePixel(img, canvas)
	WriteJpg(updateCanvas, "output.png", 90)
}

// given the side length of passport photo, returns the 4x6 canvas size
func canvasSize(imageSideLength int) (int, int) {
	canvasWidth := imageSideLength * 3
	canvasHeight := imageSideLength * 2
	return canvasWidth, canvasHeight
}

func populatePixel(img image.Image, canvas *image.RGBA) image.Image {
	w, _ := ImageSize(img)

	widthOffset := w / 3
	heightOffset := w / 2
	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			xindex := i + widthOffset
			yindex := j + heightOffset
			canvas.Set(xindex, yindex, img.At(i, j))
		}
	}

	newWidthOffset := widthOffset*2 + w
	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			xindex := i + newWidthOffset
			yindex := j + heightOffset
			canvas.Set(xindex, yindex, img.At(i, j))
		}
	}
	return canvas
}
