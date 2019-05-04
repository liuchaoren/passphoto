package passphoto

import (
	"math"
	"testing"

	"github.com/liuchaoren/passphoto/common"
)

func TestGenCanvas(t *testing.T) {
	width := 100
	height := 100
	canvas := GetCanvas(width, height)
	r := canvas.Bounds()
	if r.Dx() != width {
		t.Error("width of canvas is wrong")
	}
	if r.Dy() != height {
		t.Error("height of canvas is wrong")
	}
}

func TestCaculateHowToCrop(t *testing.T) {
	tolerance := 1e-2
	originalHeadTop, originalHeadBottom, originalHeadCenter := 0.3, 0.54, 0.6
	originalWidth, originalHeight := 1000, 1200

	cropConfig := common.ReadConfig(cropConfigFileName)
	headTopConfig, headBottomConfig := cropConfig["head_top"].(float64), cropConfig["head_bottom"].(float64)

	leftPixel, rightPixel, topPixel, bottomPixel := calculateHowToCrop(
		originalHeadTop, originalHeadBottom, originalHeadCenter, originalWidth, originalHeight)

	finalWidth, finalHeight := bottomPixel-topPixel, rightPixel-leftPixel
	if finalWidth != finalHeight {
		t.Error("The cropped image is not a square")
	}
	if math.Abs((float64(originalHeight)*originalHeadTop-float64(topPixel))/float64(finalWidth)-headTopConfig) > tolerance {
		t.Error("After crop, the haed top is not at the right position")
	}
	if math.Abs((float64(originalHeight)*0.54-float64(topPixel))/float64(finalWidth)-headBottomConfig) > tolerance {
		t.Error("After crop, the haed bottom is not at the right position")
	}
	if math.Abs(float64(originalWidth)*originalHeadCenter-float64(leftPixel)-float64(rightPixel)+float64(originalWidth)*originalHeadCenter) > 2 {
		t.Error("After crop, the head is not at the center")
	}
}
