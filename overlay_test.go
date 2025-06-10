package imgconv

import (
	"image"
	"testing"
)

func TestOverlay(t *testing.T) {
	opts := []*OverlayOption{
		&OverlayOption{
			Position: image.Point{X: 0, Y: 0},
			Opacity:  0.0,
		},
		&OverlayOption{
			Position: image.Point{X: 100, Y: 0},
			Opacity:  0.5,
		},
		&OverlayOption{
			Position: image.Point{X: 0, Y: 100},
			Opacity:  1.0,
		},
	}
	bg, err := Open("testdata/body0000.tif")
	if err != nil {
		t.Fatal(err)
	}
	over, err := Open("testdata/soft0000.tif")
	if err != nil {
		t.Fatal(err)
	}
	for _, opt := range opts {
		img := Overlay(bg, over, opt)
	}
}
