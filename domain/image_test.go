package domain

import (
	"github.com/stretchr/testify/assert"
    "os"
    "path/filepath"
    "testing"
)

func TestWrite(t *testing.T) {
    imageDir := filepath.Join(os.TempDir(), "image_dir_test")
    defer os.RemoveAll(imageDir)
    image := Image{directory: imageDir}

    Write(image, []byte{0xFF, 0xD8, 0xFF, 0x00, 0xFF, 0xD9})
    _, err := os.Stat(imageDir)
    assert.True(t, err == nil)
}