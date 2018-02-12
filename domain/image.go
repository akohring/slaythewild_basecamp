package domain

import (
	"io/ioutil"
	"path/filepath"
	"time"
	"os"
)

// Image structure
type Image struct {
	directory string
	file string
	url string
	displayTime string
}

const (
	fileExtension = ".jpg"
)

// Write the image bytes to disk under the specified directory
func Write(directory string, bytes []byte) Image {
	os.MkdirAll(directory, os.ModePerm)
	now := time.Now().UTC()
	fileName := now.Format(time.RFC3339) + fileExtension
	displayTime := now.Format("Mon Jan 02 2006 03:04:05 PM")
	file := filepath.Join(directory, fileName)
	url := "/image/"+fileName
	ioutil.WriteFile(file, bytes, 0644)
	return Image{directory: directory, file: file, url: url, displayTime: displayTime}
}