package controller

import (
	"errors"
	"encoding/json"
	"github.com/akohring/slaythewild_basecamp/domain"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestFirst(t *testing.T) {
	originalFirst := first
    defer func () { first = originalFirst }()
	var expectedImage = domain.Image{URL: "http://localhost/image/1"}
	first = func (directory string) (domain.Image, error) {
        return expectedImage, nil
	}
	controller := ImageController{directory: "/tmp/mock_image_dir"}
	w := httptest.NewRecorder()
	controller.First(w, nil)
	
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var actualImage = domain.Image{}
	err := json.Unmarshal(body, &actualImage)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedImage, actualImage)
}

func TestFirstError(t *testing.T) {
	originalFirst := first
    defer func () { first = originalFirst }()
	first = func (directory string) (domain.Image, error) {
        return domain.Image{}, errors.New("Image directory does not exist")
	}
	controller := ImageController{directory: "/tmp/mock_image_dir"}
	w := httptest.NewRecorder()
	controller.First(w, nil)
	
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var actualResponse = ImageError{}
	err := json.Unmarshal(body, &actualResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, ImageError{Message: "Image directory does not exist"}, actualResponse)
}