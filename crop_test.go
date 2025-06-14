package imgconv

import (
	"fmt"
	"path/filepath"
	"testing"
)

var (
	testCropImg = ``
)

func TestCrop(t *testing.T) {
	t.Skip()
	opts := []*CropOption{
		&CropOption{
			Width:  350,
			Height: 350,
			Anchor: Center,
		},
	}
	over, err := Open(testCropImg)
	if err != nil {
		t.Fatal(err)
	}
	for i, opt := range opts {
		img := Crop(over, opt)
		err := Save(filepath.Join("testdata", fmt.Sprintf("test-%d.tiff", i)), img, &FormatOption{})
		if err != nil {
			t.Fatal(err)
		}
	}
}
