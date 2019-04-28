package main

import (
	"os"
	"path/filepath"

	"github.com/liuchaoren/passphoto"
)

func main() {
	gopath := os.Getenv("GOPATH")
	inputImagePath := filepath.Join(gopath,
		"src/github.com/liuchaoren/passphoto/test_data/test.JPG")

	passphoto.MergePhoto(inputImagePath)

}
