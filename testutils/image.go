package testutils

import (
	"testing"
	"image"
)


func ExpectImageBoundsMatches(t *testing.T, image image.Image, minx, miny, maxx, maxy int) {
	rect := image.Bounds()
	if rect.Min.X != minx {
		t.Errorf("Image Min.X was %d, expected %d", rect.Min.X, minx)
	}

	if rect.Min.X != minx {
		t.Errorf("Image Min.Y was %d, expected %d", rect.Min.Y, miny)
	}

	if rect.Min.X != minx {
		t.Errorf("Image Max.X was %d, expected %d", rect.Max.X, maxx)
	}

	if rect.Min.X != minx {
		t.Errorf("Image Max.Y was %d, expected %d", rect.Max.Y, maxy)
	}
}
