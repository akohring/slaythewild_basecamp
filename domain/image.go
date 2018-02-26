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
	return getImageByIndex(directory, 0)
}

// Last image within a directory
func Last(directory string) (Image, error) {
	return getImageByIndex(directory, -1)
}

// Next image within a directory
func Next(directory string, currentImage string) (Image, error) {
	return getAdjacentImage(directory, currentImage, 1)
}

// Previous image within a directory
func Previous(directory string, currentImage string) (Image, error) {
	return getAdjacentImage(directory, currentImage, -1)
}

func newImage(directory string, fileName string) Image {
	file := filepath.Join(directory, fileName)
	url := "/image/"+fileName
	t, _ := time.Parse(time.RFC3339, fileName[0:len(fileName)-len(fileExtension)])
	displayTime := t.Format("Mon Jan 02 2006 03:04:05 PM")
	return Image{directory: directory, file: file, url: url, displayTime: displayTime}
}

func getImageByIndex(directory string, index int) (Image, error) {
	files, err := getSortedFiles(directory)
	if(err != nil) {
		return Image{}, err
	}
	var file os.FileInfo
	if index < 0 {
		file = files[len(files) + index]
	} else {
		file = files[index]
	}
	return newImage(directory, file.Name()), nil
}

func getAdjacentImage(directory string, currentImage string, offset int) (Image, error) {
	files, err := getSortedFiles(directory)
	if(err != nil) {
		return Image{}, err
	}
	index, err := getIndexByFile(files, currentImage)
	if(err != nil) {
		return Image{}, err
	}
	adjacentIndex := index + offset
	if(adjacentIndex == len(files)) {
		adjacentIndex = 0
	}
	return getImageByIndex(directory, adjacentIndex)
}

func getIndexByFile(files []os.FileInfo, fileName string) (int, error) {
	for index, file := range files {
		if(file.Name() == fileName) {
			return index, nil
		}
	}
	return -1, fmt.Errorf("File not found; fileName=%s", fileName)
}

func getSortedFiles(directory string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
       return files, err
	}
	if(len(files) == 0) {
		return files, fmt.Errorf("Image directory is empty; imageDir=%s", directory)
	}
	sort.Slice(files, func(i,j int) bool{
		return files[i].Name() < files[j].Name()
	})
	return files, err
}