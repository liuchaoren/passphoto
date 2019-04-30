package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// GoPath gets the GOPATH env
func GoPath() string {
	return os.Getenv("GOPATH")
}

// ConfigFolder returns config folder path
func ConfigFolder() string {
	gopath := GoPath()
	return filepath.Join(gopath, "src", "github.com",
		"liuchaoren", "passphoto", "config")
}

// ReadConfig reads the json config and return it as a map
func ReadConfig(config string) map[string]interface{} {
	cropConfigPath := filepath.Join(ConfigFolder(), config)
	cropConfigFile, err := os.Open(cropConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	defer cropConfigFile.Close()

	byteValue, _ := ioutil.ReadAll(cropConfigFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	return result
}

// ExportFolder returns the folder to export file
func ExportFolder() string {

	gopath := GoPath()
	imagePath := filepath.Join(gopath, "src", "github.com",
		"liuchaoren", "passphoto", "export_data")
	return imagePath
}
