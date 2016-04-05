package image_test

import (
	"testing"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/bcokert/render-cloud/image"
	"github.com/bcokert/render-cloud/testutils"
)

func TestColorsToNRGBAImageEmpty(t *testing.T) {
	colors := []colorful.Color{}

	result := image.ColorsToNRGBAImage(colors, 0, 0)

	testutils.ExpectImageBoundsMatches(t, result, 0, 0, 0, 0)
	testutils.ExpectUint8SlicesEqual(t, result.Pix, []uint8{})
}

func TestColorsToNRGBAImage(t *testing.T) {
	colors := []colorful.Color{
		colorful.Color{1,1,1},
		colorful.Color{0,0,0},
		colorful.Color{0,0,1},
		colorful.Color{0,1,0},
		colorful.Color{1,0,0},
		colorful.Color{0.4,0.23,0.66}, // 102, 58.65, 168.3
	}

	result := image.ColorsToNRGBAImage(colors, 3, 2)

	testutils.ExpectImageBoundsMatches(t, result, 0, 0, 3, 2)
	testutils.ExpectUint8SlicesEqual(t, result.Pix, []uint8{
		255, 255, 255, 255,
		0, 0, 0, 255,
		0, 0, 255, 255,
		0, 255, 0, 255,
		255, 0, 0, 255,
		102, 58, 168, 255,
	})
}
