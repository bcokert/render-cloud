package image

import (
	"github.com/lucasb-eyer/go-colorful"
	"image"
	"image/color"
	"image/png"
	"os"
)

const COLOR_FLOAT_FACTOR float64 = 255.0

type ImageWriter interface {
	WriteImage([]colorful.Color, uint, uint, string) error
}

type PNGImageWriter struct{}

func (*PNGImageWriter) WriteImage(pixels []colorful.Color, width, height uint, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	image := ColorsToNRGBAImage(pixels, width, height)

	if err := png.Encode(f, image); err != nil {
		return err
	}

	return nil
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
