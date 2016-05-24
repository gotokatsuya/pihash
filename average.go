package pihash

import (
	"image"
	"math"
)

type Average struct {
	Size int
}

func NewAverage() *Average {
	return &Average{
		Size: 8,
	}
}

func (a *Average) Hash(src image.Image) uint8 {
	src = ResizeImage(src, a.Size, a.Size)
	srcBounds := src.Bounds()
	maxY := srcBounds.Max.Y
	maxX := srcBounds.Max.X
	pixels := make([]uint8, maxY*maxX)
	var sumPixels uint8
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			r, g, b, _ := src.At(j, i).RGBA()
			pixel := uint8(math.Floor(float64((r + g + b) / 3)))
			pixels = append(pixels, pixel)
			sumPixels += pixel
		}
	}
	average := uint8(math.Floor(float64(sumPixels / (uint8(maxY * maxX)))))
	var (
		hash uint8
		one  uint8 = 1
	)
	for _, pixel := range pixels {
		if pixel > average {
			hash |= one
		}
		one = one << 1
	}
	return hash
}
