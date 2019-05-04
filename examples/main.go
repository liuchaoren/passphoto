package main

import (
	"path/filepath"

	"github.com/liuchaoren/passphoto"
	"github.com/liuchaoren/passphoto/common"
)

func main() {
	originalHeadTop, originalHeadBottom, originalCenterFromLeft := common.MeasureImageOnScreen()
	inputImagePath := filepath.Join(common.TestDataFolder(), "original.jpg")
	croppedImageName := filepath.Join(common.ExportFolder(), "test_crop.jpg")
	readyToPrintImageName := filepath.Join(common.ExportFolder(), "ready_to_print.jpg")
	img := passphoto.ReadImage(inputImagePath)
	passphoto.CropAndSaveJpg(img, originalHeadTop, originalHeadBottom, originalCenterFromLeft, croppedImageName, 100)
	passphoto.OutputPhotoForPrint(croppedImageName, readyToPrintImageName, 100)
}
