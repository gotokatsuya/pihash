package pihash

import "image"

type Average struct {
	Size int
}

func NewAverage() *Average {
	return &Average{
		Size: 8,
	}
}

func (a *Average) Hash(src image.Image) uint64 {
	src = GetResizedGrayscaledImage(src, a.Size, a.Size)
	srcBounds := src.Bounds()
	maxY := srcBounds.Max.Y
	maxX := srcBounds.Max.X
	var (
		pixels    []uint64
		sumPixels uint64
	)
	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			pixel := wrapSumPixels(src.At(j, i).RGBA())
			pixels = append(pixels, pixel)
			sumPixels += pixel
		}
	}
	average := uint64(sumPixels / uint64(maxY*maxX))
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
