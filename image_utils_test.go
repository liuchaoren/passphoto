package passphoto

import (
	"os"
	"path/filepath"
	"testing"
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

func TestReadJpg(t *testing.T) {
	// gopath := "/home/chaorenkindle/go"
	gopath := os.Getenv("GOPATH")
	testImagePath := filepath.Join(gopath,
		"src/github.com/liuchaoren/passphoto/test_data/test.JPG")
	ReadJpg(testImagePath)
}
