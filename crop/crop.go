package main

import (
	"flag"

	"github.com/liuchaoren/passphoto"
)

func main() {

	headTop := flag.Float64("head-top", -1., "The percent of head-top from the image top.")
	headBottom := flag.Float64("head-bottom", -1, "The percent of the head bottom (chin) from the image top.")
	centerFromLeft := flag.Float64("center-from-left", -1, "The percent of the head center from the image left.")
	image_path := flag.String("iamge-path", "", "The path of the image.")
	flag.Parse()

	img := passphoto.ReadJpg(*image_path)
	imageCropped := passphoto.Crop(img, *headTop, *headBottom, *centerFromLeft)
	passphoto.WriteJpg(imageCropped, "/home/chaorenkindle/output.jpg")

}
