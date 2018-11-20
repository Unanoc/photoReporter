package image

import (
	"io/ioutil"
	"log"
	"strings"
)

// Image contains only two fields: full path to image and name of image without format.
type Image struct {
	Path string
	Name string
}

// GetCorrectNameAndPath accepts a path to dir and image name with format then gets Full Path to Image and correct image name without format.
func (i *Image) GetCorrectNameAndPath(pathToDir, imageName string) {
	index := strings.LastIndexAny(imageName, ".") + 1 // finding substring .format, e.g. "*.jpg"
	newImageName := imageName[:index]
	newFullPathToImage := pathToDir + "/" + imageName
	i.Path = newFullPathToImage
	i.Name = newImageName
}

// GetPhotos returns a slice of Image.
func GetPhotos(path string) []Image {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	images := make([]Image, 0)
	for _, f := range files {
		if f.Name() != ".DS_Store" {
			image := Image{}
			image.GetCorrectNameAndPath(path, f.Name())
			images = append(images, image)
		}
	}

	return images
}
