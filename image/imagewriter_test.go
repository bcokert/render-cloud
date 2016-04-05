package image_test

import (
	"github.com/bcokert/render-cloud/image"
	"github.com/bcokert/render-cloud/testutils"
	"github.com/lucasb-eyer/go-colorful"
	stdImage "image"
	"io"
	"testing"
	"errors"
	"os"
)

func TestColorsToNRGBAImageEmpty(t *testing.T) {
	colors := []colorful.Color{}

	result := image.ColorsToNRGBAImage(colors, 0, 0)

	testutils.ExpectImageBoundsMatches(t, result, 0, 0, 0, 0)
	testutils.ExpectUint8SlicesEqual(t, result.Pix, []uint8{})
}

func TestColorsToNRGBAImage(t *testing.T) {
	colors := []colorful.Color{
		colorful.Color{1, 1, 1},
		colorful.Color{0, 0, 0},
		colorful.Color{0, 0, 1},
		colorful.Color{0, 1, 0},
		colorful.Color{1, 0, 0},
		colorful.Color{0.4, 0.23, 0.66}, // 102, 58.65, 168.3
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

func mockPNGEncoderNoop(w io.Writer, i stdImage.Image) error {
	return nil
}

func mockPNGEncoderFails(w io.Writer, i stdImage.Image) error {
	return errors.New("I dun broke")
}

func TestPNGImageWriterWriteImageErrorEmptyImage(t *testing.T) {
	colors := []colorful.Color{}
	pngWriter := image.PNGImageWriter{}
	err := pngWriter.WriteImage(mockPNGEncoderNoop, colors, 0, 0, "")

	if err == nil {
		t.Errorf("PNGImageWriter.WriteImage did not fail when given an empty image")
	}

	if err.Error() != "Cannot write an image with zero pixels" {
		t.Errorf("PNGImageWriter.WriteImage did not give the correct error message when an empty image was given")
	}
}

func TestPNGImageWriterWriteImageErrorZeroHeightWidth(t *testing.T) {
	colors := []colorful.Color{
		colorful.Color{1, 1, 1},
	}
	pngWriter := image.PNGImageWriter{}
	err1 := pngWriter.WriteImage(mockPNGEncoderNoop, colors, 2, 0, "")
	err2 := pngWriter.WriteImage(mockPNGEncoderNoop, colors, 0, 5, "")

	if err1 == nil || err2 == nil {
		t.Errorf("PNGImageWriter.WriteImage did not fail when given a zero width or height")
	}

	if err1.Error() != "Cannot write an image with zero width or height" || err2.Error() != "Cannot write an image with zero width or height" {
		t.Errorf("PNGImageWriter.WriteImage did not give the correct error message when given a zero width or height")
	}
}

func TestPNGImageWriterWriteImageErrorDimensionsDontMatchSize(t *testing.T) {
	colors := []colorful.Color{
		colorful.Color{1, 1, 1},
		colorful.Color{1, 1, 1},
	}
	pngWriter := image.PNGImageWriter{}
	err1 := pngWriter.WriteImage(mockPNGEncoderNoop, colors, 2, 3, "")
	err2 := pngWriter.WriteImage(mockPNGEncoderNoop, colors, 1, 1, "")

	if err1 == nil || err2 == nil {
		t.Errorf("PNGImageWriter.WriteImage did not fail when given an image whos dimensions did not match the number of pixels")
	}

	if err1.Error() != "Cannot write a 2x3 Image when given 2 pixels, require 6" {
		t.Errorf("PNGImageWriter.WriteImage did not give the correct error message when pixels was less than the dimensions")
	}

	if err2.Error() != "Cannot write a 1x1 Image when given 2 pixels, require 1" {
		t.Errorf("PNGImageWriter.WriteImage did not give the correct error message when pixels was more than the dimensions")
	}
}

func TestPNGImageWriterWriteImageFileOpenFails(t *testing.T) {
	colors := []colorful.Color{
		colorful.Color{1, 1, 1},
	}
	pngWriter := image.PNGImageWriter{}
	err := pngWriter.WriteImage(mockPNGEncoderNoop, colors, 1, 1, "")

	if err == nil {
		t.Errorf("PNGImageWriter.WriteImage did not fail when given an empty file name")
	}

	if err.Error() != "Error opening file: open : no such file or directory" {
		t.Errorf("PNGImageWriter.WriteImage did not give the correct error message when an empty file was given")
	}
}

func TestPNGImageWriterWriteEncodingFails(t *testing.T) {
	colors := []colorful.Color{
		colorful.Color{1, 1, 1},
	}
	pngWriter := image.PNGImageWriter{}

	err := pngWriter.WriteImage(mockPNGEncoderFails, colors, 1, 1, "./testfile.png")
	defer func () {
		err := os.Remove("./testfile.png")
		if err != nil {
			t.Errorf("Failed to remove temporary file 'testfile.png' created during test")
		}
	}()

	if err == nil {
		t.Errorf("PNGImageWriter.WriteImage did not fail when encoding to png failed")
	}

	if err.Error() != "Error encoding png: I dun broke" {
		t.Errorf("PNGImageWriter.WriteImage did not give the correct error message when encoding to png failed")
	}
}

func TestPNGImageWriterWriteImage(t *testing.T) {
	colors := []colorful.Color{
		colorful.Color{1, 1, 1},
		colorful.Color{1, 1, 1},
		colorful.Color{1, 1, 1},
		colorful.Color{1, 1, 1},
		colorful.Color{0, 0, 1},
		colorful.Color{0, 1, 0},
		colorful.Color{1, 0, 0},
		colorful.Color{0, 0, 0},
	}
	pngWriter := image.PNGImageWriter{}

	err := pngWriter.WriteImage(image.DefaultPNGEncoder, colors, 4, 2, "./testfile2.png")
	defer func () {
		err := os.Remove("./testfile2.png")
		if err != nil {
			t.Errorf("Failed to remove temporary file 'testfile2.png' created during test")
		}
	}()

	if err != nil {
		t.Errorf("PNGImageWriter.WriteImage did not succeed with legal inputs")
	}
}
