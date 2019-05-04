package common

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// MeasureImageOnScreen measure the image by clicking on image
func MeasureImageOnScreen() (float64, float64, float64) {

	fmt.Println("--- Please click the top left corner of image ---")
	robotgo.AddEvent("mleft")
	leftTopX, leftTopY := robotgo.GetMousePos()

	fmt.Println("--- Please click the bottom right corner of image ---")
	robotgo.AddEvent("mleft")
	rightBottomX, rightBottomY := robotgo.GetMousePos()

	fmt.Println("--- Please click the head top ---")
	robotgo.AddEvent("mleft")
	_, headTopY := robotgo.GetMousePos()

	fmt.Println("--- Please click the head bottom ---")
	robotgo.AddEvent("mleft")
	_, headBottomY := robotgo.GetMousePos()

	fmt.Println("--- Please click the nose ---")
	robotgo.AddEvent("mleft")
	noseX, _ := robotgo.GetMousePos()

	width, height := rightBottomX-leftTopX, rightBottomY-leftTopY

	originalHeadTop := float64(headTopY-leftTopY) / float64(height)
	originalHeadBottom := float64(headBottomY-leftTopY) / float64(height)
	originalCenterFromLeft := float64(noseX-leftTopX) / float64(width)
	return originalHeadTop, originalHeadBottom, originalCenterFromLeft
}
