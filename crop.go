package imgconv

import "image"

// CropOption is overlay option
type CropOption struct {
	Width  int
	Height int
	Anchor Anchor
}

// Crop draws an image over another at position with opacity.
func Crop(base image.Image, option *CropOption) image.Image {
	return option.do(base)
}

func (c *CropOption) do(base image.Image) image.Image {
	return cropAnchor(base, c.Width, c.Height, c.Anchor)
}
