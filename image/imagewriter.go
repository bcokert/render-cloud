package image

import (
	"errors"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
)

const COLOR_FLOAT_FACTOR float64 = 255.0

type ImageWriter interface {
	WriteImage([]colorful.Color, uint, uint, string) error
}

func ColorsToNRGBAImage(colors []colorful.Color, width, height uint) *image.NRGBA {
	ww, hh := int(width), int(height)
	image := image.NewNRGBA(image.Rectangle{image.Point{0, 0}, image.Point{ww, hh}})
	for y := 0; y < hh; y++ {
		for x := 0; x < ww; x++ {
			r := uint8(colors[y*ww+x].R * COLOR_FLOAT_FACTOR)
			g := uint8(colors[y*ww+x].G * COLOR_FLOAT_FACTOR)
			b := uint8(colors[y*ww+x].B * COLOR_FLOAT_FACTOR)
			image.SetNRGBA(x, y, color.NRGBA{r, g, b, 255})
		}
	}
	return image
}

var DefaultPNGEncoder = png.Encode

type PNGImageWriter struct{}

func (*PNGImageWriter) WriteImage(encoder func(io.Writer, image.Image) error, pixels []colorful.Color, width, height uint, filename string) error {
	numPixels := len(pixels)
	if numPixels == 0 {
		return errors.New("Cannot write an image with zero pixels")
	}

	if width == 0 || height == 0 {
		return errors.New("Cannot write an image with zero width or height")
	}

	if width*height != uint(numPixels) {
		return errors.New(fmt.Sprintf("Cannot write a %dx%d Image when given %d pixels, require %d", width, height, numPixels, width*height))
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return errors.New("Error opening file: " + err.Error())
	}

	image := ColorsToNRGBAImage(pixels, width, height)

	if err := encoder(f, image); err != nil {
		return errors.New("Error encoding png: " + err.Error())
	}

	return nil
}
