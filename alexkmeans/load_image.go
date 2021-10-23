package alexkmeans

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func LoadImage(imagePath string) (image.Image, error) {
	f, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if strings.HasSuffix(imagePath, "png") {
		image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	}

	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

	image, _, err := image.Decode(f)
	return image, err
}
