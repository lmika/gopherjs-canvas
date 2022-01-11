package imgprocess

import (
	"bytes"
	"github.com/disintegration/imaging"
	"image"
	"image/png"
)

func ReadImage(bs []byte) (image.Image, error) {
	rawPng, err := png.Decode(bytes.NewReader(bs))
	if err != nil {
		return nil, err
	}

	resizeImg := imaging.Resize(rawPng, 300, 0, imaging.Gaussian)
	// resizeImg := resize.Resize(300, 0, rawPng, resize.NearestNeighbor)
	return resizeImg, nil
}
