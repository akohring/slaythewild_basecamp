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
    imageDir := mockImageDir()
    defer os.RemoveAll(imageDir)
    seedImageDir(imageDir)

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
    imageDir := mockImageDir()
    defer os.RemoveAll(imageDir)

    _, err := First(imageDir)
    assert.Equal(t, fmt.Sprintf("Image directory is empty; imageDir=%s", imageDir), err.Error())
}

func TestLast(t *testing.T) {
    imageDir := mockImageDir()
    defer os.RemoveAll(imageDir)
    seedImageDir(imageDir)

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
    imageDir := mockImageDir()
    defer os.RemoveAll(imageDir)

    _, err := Last(imageDir)
    assert.Equal(t, fmt.Sprintf("Image directory is empty; imageDir=%s", imageDir), err.Error())
}

func TestNext(t *testing.T) {
    imageDir := mockImageDir()
    defer os.RemoveAll(imageDir)
    seedImageDir(imageDir)

    image, err := Next(imageDir, "2016.jpg")
    assert.Nil(t, err)
    assert.Equal(t, filepath.Join(imageDir, "2017.jpg"), image.file)
}

func TestNextAtEndOfList(t *testing.T) {
    imageDir := mockImageDir()
    defer os.RemoveAll(imageDir)
    seedImageDir(imageDir)

    image, err := Next(imageDir, "2017.jpg")
    assert.Nil(t, err)
    assert.Equal(t, filepath.Join(imageDir, "2014.jpg"), image.file)
}

func TestNextErrorWhenDirectoryDoesNotExist(t *testing.T) {
    imageDir := filepath.Join(os.TempDir(), "image_dir_test")
    _, err := Next(imageDir, "2016.jpg")
    assert.Equal(t, fmt.Sprintf("open %s: no such file or directory", imageDir), err.Error())
}

func TestNextErrorWhenDirectoryIsEmpty(t *testing.T) {
    imageDir := mockImageDir()
    defer os.RemoveAll(imageDir)

    _, err := Next(imageDir, "2016.jpg")
    assert.Equal(t, fmt.Sprintf("Image directory is empty; imageDir=%s", imageDir), err.Error())
}

func TestNextErrorWhenFileNotFound(t *testing.T) {
    imageDir := mockImageDir()
    defer os.RemoveAll(imageDir)
    seedImageDir(imageDir)

    _, err := Next(imageDir, "doesnotexist.jpg")
    assert.Equal(t, "File not found; fileName=doesnotexist.jpg", err.Error())
}

func TestPrevious(t *testing.T) {
    imageDir := mockImageDir()
    defer os.RemoveAll(imageDir)
    seedImageDir(imageDir)

    image, err := Previous(imageDir, "2016.jpg")
    assert.Nil(t, err)
    assert.Equal(t, filepath.Join(imageDir, "2015.jpg"), image.file)
}

func TestPreviousFromBeginningOfList(t *testing.T) {
    imageDir := mockImageDir()
    defer os.RemoveAll(imageDir)
    seedImageDir(imageDir)

    image, err := Previous(imageDir, "2014.jpg")
    assert.Nil(t, err)
    assert.Equal(t, filepath.Join(imageDir, "2017.jpg"), image.file)
}

func TestPreviousErrorWhenDirectoryDoesNotExist(t *testing.T) {
    imageDir := filepath.Join(os.TempDir(), "image_dir_test")
    _, err := Previous(imageDir, "2016.jpg")
    assert.Equal(t, fmt.Sprintf("open %s: no such file or directory", imageDir), err.Error())
}

func TestPreviousErrorWhenDirectoryIsEmpty(t *testing.T) {
    imageDir := mockImageDir()
    defer os.RemoveAll(imageDir)

    _, err := Previous(imageDir, "2016.jpg")
    assert.Equal(t, fmt.Sprintf("Image directory is empty; imageDir=%s", imageDir), err.Error())
}

func TestPreviousErrorWhenFileNotFound(t *testing.T) {
    imageDir := mockImageDir()
    defer os.RemoveAll(imageDir)
    seedImageDir(imageDir)

    _, err := Previous(imageDir, "doesnotexist.jpg")
    assert.Equal(t, "File not found; fileName=doesnotexist.jpg", err.Error())
}

func mockImageDir() string {
    imageDir := filepath.Join(os.TempDir(), "image_dir_test")
    os.MkdirAll(imageDir, os.ModePerm)
    return imageDir
}

func seedImageDir(imageDir string) {
    bytes := []byte{0xFF, 0xD8, 0xFF, 0x00, 0xFF, 0xD9}
    files := []string{"2015.jpg", "2016.jpg", "2017.jpg", "2014.jpg"}
    for _, file := range files {
        ioutil.WriteFile(filepath.Join(imageDir, file), bytes, 0644)
    }
}