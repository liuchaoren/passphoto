package passphoto

import (
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"path/filepath"

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

func WriteJpg(img *image.RGBA, path string) {
	gopath := os.Getenv("GOPATH")
	imagePath := filepath.Join(gopath, "src/github.com/liuchaoren/passphoto/export_data", path)

	f, err := os.Create(imagePath)
	if err != nil {
		log.Fatal("cannot write file")
	}
	png.Encode(f, img)
}

func goPath() string {
	return os.Getenv("GOPATH")
}

func cropConfig() map[string]interface{} {
	gopath := goPath()
	configPath := filepath.Join(gopath, "src", "github.com", "liuchaoren", "passphoto", "config")
	cropConfigPath := filepath.Join(configPath, "crop.json")

	cropConfigFile, err := os.Open(cropConfigPath)
	if err != nil {
		log.Fatal("cannot open crop config file")
	}
	defer cropConfigFile.Close()

	byteValue, _ := ioutil.ReadAll(cropConfigFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	return result
}

func Crop(img image.Image, originalHeadTop, orginalHeadBottom, centerFromLeft float64) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	cropConfig := cropConfig()

}

func widthAndHeight(img image.Image) (int, int) {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	return width, height
}

func crop(img image.Image, leftFromLeft, rightFromLeft, topFromTop, bottomFromTop float64) {
	w, h := widthAndHeight(img)

}
