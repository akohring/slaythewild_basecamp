package domain

import (
    "fmt"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
    "os"
    "path/filepath"
    "testing"
)

func TestWrite(t *testing.T) {
    imageDir := filepath.Join(os.TempDir(), "image_dir_test")
    defer os.RemoveAll(imageDir)

    image := Write(imageDir, []byte{0xFF, 0xD8, 0xFF, 0x00, 0xFF, 0xD9})
    assert.Equal(t, imageDir, image.directory)
    assert.Equal(t, "/image/"+filepath.Base(image.file), image.url)
    assert.NotNil(t, image.displayTime)
    
    _, err := os.Stat(imageDir)
    assert.Nil(t, err)

    _, err = os.Stat(image.file)
    assert.Nil(t, err)
}

func TestFirst(t *testing.T) {
    imageDir := filepath.Join(os.TempDir(), "image_dir_test")
    os.MkdirAll(imageDir, os.ModePerm)
    defer os.RemoveAll(imageDir)

    bytes := []byte{0xFF, 0xD8, 0xFF, 0x00, 0xFF, 0xD9}
    files := []string{"2015.jpg", "2016.jpg", "2017.jpg", "2014.jpg"}
    for _, file := range files {
        ioutil.WriteFile(filepath.Join(imageDir, file), bytes, 0644)
    }

    image, err := First(imageDir)
    assert.Nil(t, err)
    assert.Equal(t, filepath.Join(imageDir, "2014.jpg"), image.file)
}

func TestFirstErrorWhenDirectoryDoesNotExist(t *testing.T) {
    imageDir := filepath.Join(os.TempDir(), "image_dir_test")
    _, err := First(imageDir)
    assert.Equal(t, fmt.Sprintf("open %s: no such file or directory", imageDir), err.Error())
}

func TestFirstErrorWhenDirectoryIsEmpty(t *testing.T) {
    imageDir := filepath.Join(os.TempDir(), "image_dir_test")
    os.MkdirAll(imageDir, os.ModePerm)
    defer os.RemoveAll(imageDir)

    _, err := First(imageDir)
    assert.Equal(t, fmt.Sprintf("Image directory is empty; imageDir=%s", imageDir), err.Error())
}

func TestLast(t *testing.T) {
    imageDir := filepath.Join(os.TempDir(), "image_dir_test")
    os.MkdirAll(imageDir, os.ModePerm)
    defer os.RemoveAll(imageDir)

    bytes := []byte{0xFF, 0xD8, 0xFF, 0x00, 0xFF, 0xD9}
    files := []string{"2015.jpg", "2016.jpg", "2017.jpg", "2014.jpg"}
    for _, file := range files {
        ioutil.WriteFile(filepath.Join(imageDir, file), bytes, 0644)
    }

    image, err := Last(imageDir)
    assert.Nil(t, err)
    assert.Equal(t, filepath.Join(imageDir, "2017.jpg"), image.file)
}

func TestLastErrorWhenDirectoryDoesNotExist(t *testing.T) {
    imageDir := filepath.Join(os.TempDir(), "image_dir_test")
    _, err := Last(imageDir)
    assert.Equal(t, fmt.Sprintf("open %s: no such file or directory", imageDir), err.Error())
}

func TestLastErrorWhenDirectoryIsEmpty(t *testing.T) {
    imageDir := filepath.Join(os.TempDir(), "image_dir_test")
    os.MkdirAll(imageDir, os.ModePerm)
    defer os.RemoveAll(imageDir)

    _, err := Last(imageDir)
    assert.Equal(t, fmt.Sprintf("Image directory is empty; imageDir=%s", imageDir), err.Error())
}