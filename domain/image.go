package domain

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"
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
	now := time.Now().UTC()
	fileName := now.Format(time.RFC3339) + fileExtension
	image := newImage(directory, fileName)
	os.MkdirAll(directory, os.ModePerm)
	ioutil.WriteFile(image.file, bytes, 0644)
	return image
}

// First image within a directory
func First(directory string) (Image, error) {
	return getFileByIndex(directory, 0)
}

func newImage(directory string, fileName string) Image {
	file := filepath.Join(directory, fileName)
	url := "/image/"+fileName
	t, _ := time.Parse(time.RFC3339, fileName[0:len(fileName)-len(fileExtension)])
	displayTime := t.Format("Mon Jan 02 2006 03:04:05 PM")
	return Image{directory: directory, file: file, url: url, displayTime: displayTime}
}

func getFileByIndex(directory string, index int) (Image, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
       return Image{}, err
	}
	if(len(files) == 0) {
		return Image{}, fmt.Errorf("Image directory is empty; imageDir=%s", directory)
	}
	sort.Slice(files, func(i,j int) bool{
		return files[i].Name() < files[j].Name()
	})
	return newImage(directory, files[index].Name()), nil
}