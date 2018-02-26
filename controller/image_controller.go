package controller

import (
    "github.com/akohring/slaythewild_basecamp/domain"
    "encoding/json"
    "net/http"
)

// ImageController structure
type ImageController struct {
    directory string
}

// ImageError structure
type ImageError struct {
	Message string `json:"msg"`
}

var first = domain.First

// First image in the directory
func (c *ImageController) First(w http.ResponseWriter, r *http.Request) {
    image, err := first(c.directory)
    if(err != nil) {
        json.NewEncoder(w).Encode(ImageError{Message: err.Error()})
    } else {
        json.NewEncoder(w).Encode(image)
    }
}