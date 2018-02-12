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

    image := Write(imageDir, []byte{0xFF, 0xD8, 0xFF, 0x00, 0xFF, 0xD9})
    assert.Equal(t, imageDir, image.directory)
    assert.Equal(t, "/image/"+filepath.Base(image.file), image.url)
    assert.NotNil(t, image.displayTime)
    
    _, err := os.Stat(imageDir)
    assert.Nil(t, err)

    _, err = os.Stat(image.file)
    assert.Nil(t, err)
}