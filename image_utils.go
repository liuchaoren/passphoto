package passphoto

import (
	"image"
	"image/color"
	"image/jpeg"

	// need jpeg to decode
	_ "image/jpeg"
	"log"
	"os"

	"github.com/liuchaoren/passphoto/common"
)

const (
	cropConfigFileName string = "crop.json"
	jgpQuality         int    = 80
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

// ReadImage reads a jpg image
func ReadImage(path string) *image.RGBA {

	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	mRGBA := image.NewRGBA(m.Bounds())
	width, height := mRGBA.Bounds().Dx(), mRGBA.Bounds().Dy()
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			mRGBA.Set(i, j, m.At(i, j))
		}
	}
	return mRGBA
}

// WriteJpg writes an image to jpg file
func WriteJpg(img image.Image, path string, quality int) {

	f, err := os.Create(path)
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

// caculate how to cut the four edges.
//
// Args:
//	originalHeadTop: percentage of the position of head top from image top edge
//	originalHeadBottom: percentage of the position of head bottom from image top edge
//	originalCenterFromLeft: percentage of the position of head center from image left edge
//	originalWidth, originalHeight: width, height of the original image.
//
// Returns:
//	cutLeft, cutRight: position to cut on left/right side, counting from the left edge.
//	cutTop, cutBottom: position to cut on top/bottom side, counting from the top edge.
func calculateHowToCrop(originalHeadTop, originalHeadBottom,
	originalCenterFromLeft float64, originalWidth, originalHeight int) (
	int, int, int, int) {

	cropConfig := common.ReadConfig(cropConfigFileName)
	headTop := cropConfig["head_top"].(float64)
	headBottom := cropConfig["head_bottom"].(float64)

	cutTop := (originalHeadTop*headBottom - originalHeadBottom*headTop) / (headBottom - headTop)
	cutBottom := (originalHeadTop + (headTop-1)*cutTop) / headTop

	squareSideLength := int((cutBottom - cutTop) * float64(originalHeight))
	cutLeft := originalCenterFromLeft - (float64(squareSideLength) / float64(2) / float64(originalWidth))
	cutRight := originalCenterFromLeft + (float64(squareSideLength) / float64(2) / float64(originalWidth))

	leftPixel := int(float64(originalWidth) * cutLeft)
	rightPixel := int(float64(originalWidth) * cutRight)
	topPixel := int(float64(originalHeight) * cutTop)
	bottomPixel := int(float64(originalHeight) * cutBottom)

	tmpWidth := rightPixel - leftPixel
	tmpHeight := bottomPixel - topPixel

	if tmpWidth > tmpHeight {
		squareSideLength = tmpWidth
	} else {
		squareSideLength = tmpHeight
	}

	rightPixel = leftPixel + squareSideLength
	bottomPixel = topPixel + squareSideLength

	return leftPixel, rightPixel, topPixel, bottomPixel
}

// Crop crops image given the vertical position of head top, head bottom
// and horiental position of head cener
//
// Args:
// 	img: image to cut
//	originalHeadTop: percentage of the position of head top from image top edge
//	originalHeadBottom: percentage of the position of head bottom from image top edge
//	originalCenterFromLeft: percentage of the position of head center from image left edge
func Crop(img image.Image, originalHeadTop, orginalHeadBottom,
	originalCenterFromLeft float64) image.Image {
	width, height := ImageSize(img)
	leftPixel, rightPixel, topPixel, bottomPixel := calculateHowToCrop(originalHeadTop,
		orginalHeadBottom, originalCenterFromLeft, width, height)
	return crop(img, leftPixel, rightPixel, topPixel, bottomPixel)
}

// CropAndSaveJpg crops the image and persist on disk
func CropAndSaveJpg(img image.Image, originalHeadTop, orginalHeadBottom,
	originalCenterFromLeft float64, outputPath string, quality int) {
	cropedImage := Crop(img, originalHeadTop, orginalHeadBottom, originalCenterFromLeft)
	WriteJpg(cropedImage, outputPath, quality)
}

// crop acutally does the crop and returns the cropped image
func crop(img image.Image, leftPixel, rightPixel,
	topPixel, bottomPixel int) image.Image {
	imgRGBA, ok := img.(*image.RGBA)
	if !ok {
		log.Fatal("cannot convert to RGBA image")
	}

	upLeft := image.Point{leftPixel, topPixel}
	downRight := image.Point{rightPixel, bottomPixel}
	subImage := imgRGBA.SubImage(image.Rectangle{upLeft, downRight})
	return subImage
}

// OutputPhotoForPrint prepares for print on a 4x6 canvas
func OutputPhotoForPrint(inputPath, outputPath string, quality int) {

	img := ReadImage(inputPath)

	w, h := ImageSize(img)
	if w != h {
		log.Fatal("image must be a square")
	}
	canvasWidth, canvasHeight := canvasSize(w)
	canvas := GetCanvas(canvasWidth, canvasHeight)

	updateCanvas := populatePixel(img, canvas)
	WriteJpg(updateCanvas, outputPath, quality)
}

// given the side length of passport photo, returns the 4x6 canvas size
func canvasSize(imageSideLength int) (int, int) {
	canvasWidth := imageSideLength * 3
	canvasHeight := imageSideLength * 2
	return canvasWidth, canvasHeight
}

func populatePixel(img image.Image, canvas image.Image) *image.RGBA {
	w, _ := ImageSize(img)

	widthOffset := w / 3
	heightOffset := w / 2
	canvasRGBA := canvas.(*image.RGBA)
	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			xindex := i + widthOffset
			yindex := j + heightOffset
			canvasRGBA.Set(xindex, yindex, img.At(i, j))
		}
	}

	newWidthOffset := widthOffset*2 + w
	for i := 0; i < w; i++ {
		for j := 0; j < w; j++ {
			xindex := i + newWidthOffset
			yindex := j + heightOffset
			canvasRGBA.Set(xindex, yindex, img.At(i, j))
		}
	}
	return canvasRGBA
}
