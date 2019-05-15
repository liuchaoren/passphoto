package main

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/liuchaoren/passphoto"
	"github.com/liuchaoren/passphoto/common"
)

func main() {
	inputImage := flag.String("input-image", "EMPTY", "Required path to the input image file.")
	outputImage := flag.String("output-image", "EMPTY", "Required path to write the output image file.")
	croppedImage := flag.String("cropped-image", "EMPTY", "Optional path to save the cropped image.")

	flag.Parse()
	if *inputImage == "EMPTY" || *outputImage == "EMPTY" {
		log.Fatal("-input-image and -output-image are required!")
	}
	if *croppedImage == "EMPTY" {
		tempDir, err := ioutil.TempDir("", "photo")
		if err != nil {
			log.Fatal()
		}
		*croppedImage = filepath.Join(tempDir, "cropped_image.jpg")
	}

	originalHeadTop, originalHeadBottom, originalCenterFromLeft := common.MeasureImageOnScreen()
	img := passphoto.ReadImage(*inputImage)
	passphoto.CropAndSaveJpg(img, originalHeadTop, originalHeadBottom, originalCenterFromLeft, *croppedImage, 100)
	passphoto.OutputPhotoForPrint(*croppedImage, *outputImage, 100)

}
