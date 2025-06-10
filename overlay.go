package imgconv

import "image"

// OverlayOption is overlay option
type OverlayOption struct {
	Position image.Point
	Opacity  float64
}

// Overlay draws an image over another at position with opacity.
func Overlay(bg, over image.Image, option *OverlayOption) image.Image {
	return option.do(bg, over)
}

func (o *OverlayOption) do(bg, over image.Image) image.Image {
	return overlay(bg, over, o.Position, o.Opacity)
}
