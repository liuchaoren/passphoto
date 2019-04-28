package main

import (
	"flag"
	"fmt"

	"github.com/liuchaoren/passphoto"
)

func main() {

	head_top := flag.Float64("head-top", -1., "The percent of head-top from the image top.")
	head_bottom := flag.Float64("head-bottom", -1, "The percent of the head bottom (chin) from the image top.")
	image_path := flag.String("iamge-path", "", "The path of the image.")
	flag.Parse()

	fmt.Println(*head_top)
	fmt.Println(*head_bottom)

	img := passphoto.ReadJpg(*image_path)
	passphoto.Crop(img)

}
