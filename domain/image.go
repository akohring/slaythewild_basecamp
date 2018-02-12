package domain

import (
	"os"
)

// Image structure
type Image struct {
	directory string
}

// Write will write the image bytes to disk
func Write(image Image, bytes []byte) {
	os.MkdirAll(image.directory, os.ModePerm)
}