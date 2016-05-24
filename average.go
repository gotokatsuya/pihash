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

func (a *Average) Hash(src image.Image) uint64 {
	src = ResizeImage(src, a.Size, a.Size)
	srcBounds := src.Bounds()
	maxY := srcBounds.Max.Y
	maxX := srcBounds.Max.X
	var (
		pixels    []uint64
		sumPixels uint64
	)
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			r, g, b, _ := src.At(j, i).RGBA()
			pixel := uint64(math.Floor(float64((r + g + b)) / float64(3)))
			pixels = append(pixels, pixel)
			sumPixels += pixel
		}
	}
	average := uint64(math.Floor(float64(sumPixels) / float64((maxY * maxX))))
	var (
		hash uint64
		one  uint64 = 1
	)
	for _, pixel := range pixels {
		if pixel > average {
			hash |= one
		}
		one = one << 1
	}
	return hash
}
